package taskrunner

import (
	"backend/pkg/logger"
	"context"
	"fmt"
	"time"
)

type PeriodicTask struct {
	Action Task
	Ticker *time.Ticker
	Log    logger.Logger
}

func NewPeriodicTask(action Task, ticker *time.Ticker) PeriodicTask {
	return PeriodicTask{
		Action: action,
		Ticker: ticker,
		Log:    logger.GetLogrusLogger(),
	}
}

func (t PeriodicTask) Run(ctx context.Context) {
	t.Log.Info("Running background Jobs")
	errChan := make(chan error, 3)
	statusChannel := make(chan string, 3)
	//t.Action(errChan, statusChan)

	for {
		select {
		case <-ctx.Done():
			t.Log.Info("context cancelled", t.Action.Name)
			return
		case err := <-errChan:
			t.Log.Error("err occured", err)
			//return err
		case status := <-statusChannel:
			t.Log.Info("status", status)

		case <-t.Ticker.C:
			t.Log.Info(" executing periodic the task", t.Action.Name)
			go func(errCh chan error, statusCh chan string) {
				defer func() {
					if r := recover(); r != nil {
						// Handle the panic and send the error to the error channel
						errCh <- fmt.Errorf("panic occurred: %v", r)
					}
				}()
				t.Action.Invoke(errCh, statusCh)
			}(errChan, statusChannel)

		}
	}

}

func (t PeriodicTask) GetName() string {
	return t.Action.Name
}
