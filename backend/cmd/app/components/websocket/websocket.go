package websocket

import (
	"backend/pkg/logger"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	socketio "github.com/googollee/go-socket.io"
)

type Websocket struct {
	wsserver     *socketio.Server
	Addr         string
	Name         string
	httpserver   *http.Server
	Log          logger.Logger
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func (ws *Websocket) Run(ctx context.Context) error {
	ws.initServer()
	errChan := make(chan error, 1)
	go func() {
		err := ws.start()
		if err != nil {
			errChan <- err
		}
	}()

	go func() {
		err := ws.wsserver.Serve()
		if err != nil {
			errChan <- err
		}
	}()

	for {

		select {
		case <-ctx.Done():
			ws.Log.Info("context cancelled", ws.Name)
			ws.httpserver.Shutdown(ctx)
			ws.wsserver.Close()
			return errors.New("context cancelled")
		case err := <-errChan:
			ws.Log.Info("err occured", ws.Name)
			ws.httpserver.Shutdown(ctx)
			ws.wsserver.Close()
			return err

		}

	}
}
func (ws *Websocket) initServer() {
	ws.wsserver = socketio.NewServer(nil)

	ws.wsserver.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	ws.wsserver.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})

	ws.wsserver.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})

	ws.wsserver.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	ws.wsserver.OnError("/", func(s socketio.Conn, e error) {
		// server.Remove(s.ID())
		fmt.Println("meet error:", e)
	})

	ws.wsserver.OnDisconnect("/", func(s socketio.Conn, reason string) {
		// Add the Remove session id. Fixed the connection & mem leak
		fmt.Println("closed", reason)
	})

}

func (ws *Websocket) start() error {
	ws.Log.Info("starting wsserver", ws.Name)
	ws.httpserver = &http.Server{
		Addr:         ws.Addr,
		Handler:      ws.wsserver,
		ReadTimeout:  ws.ReadTimeout,
		WriteTimeout: ws.WriteTimeout,
	}
	err := ws.httpserver.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil

}
func (ws *Websocket) GetName() string {
	return ws.Name
}
