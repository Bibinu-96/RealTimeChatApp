package server

import (
	"net/http"
	"time"

	"backend/cmd/app/service/logger"

	"backend/cmd/app/service/server/config"
)

type Server struct {
	Server       *http.Server
	Router       http.Handler
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Name         string
	log          logger.Logger
}

func NewServer(config config.ServerConfig, name string, log logger.Logger) *Server {

	return &Server{
		Router:       config.Router,
		Addr:         config.Addr,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		Name:         name,
		log:          log,
	}

}

func (s *Server) Run() error {
	return s.Start()
}
func (s *Server) GetName() string {
	return s.Name
}

func (s *Server) Start() error {
	s.log.Info("starting server", s.Name)
	s.Server = &http.Server{
		Addr:         s.Addr,
		Handler:      s.Router,
		ReadTimeout:  s.ReadTimeout,
		WriteTimeout: s.WriteTimeout,
	}
	err := s.Server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
	}
	return err

}
