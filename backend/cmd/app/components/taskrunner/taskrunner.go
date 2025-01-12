package taskrunner

import (
	"backend/internal/channels"
	"backend/pkg/logger"
	"context"
	"errors"
	"fmt"
	"sync"
)

type TaskRunner struct {
	Log  logger.Logger
	Name string
	wg   sync.WaitGroup
}

func (tr TaskRunner) Run(ctx context.Context) error {
	tr.Log.Info("Running background Jobs")
	errChan := make(chan error, 3)
	statusChannel := make(chan string, 3)
	taskChannel := channels.GetTaskChannel()
	for {

		select {
		case <-ctx.Done():
			tr.Log.Info("context cancelled", tr.Name)
			tr.wg.Wait()
			return errors.New("context cancelled")
		case err := <-errChan:
			tr.Log.Error("err occured", err)
			//return err
		case status := <-statusChannel:
			tr.Log.Info("status", status)

		case config := <-taskChannel:
			switch castedValue := config.(type) {
			case Task:
				tr.Log.Info(" executing the task")
				go func(errCh chan error, statusCh chan string) {
					defer func() {
						if r := recover(); r != nil {
							// Handle the panic and send the error to the error channel
							errCh <- fmt.Errorf("panic occurred: %v", r)
						}
					}()
					castedValue.Invoke(errCh, statusCh)
				}(errChan, statusChannel)
			case PeriodicTask:
				tr.Log.Info(" Intialising Periodic task", castedValue.Action.Name)
				// create derived context
				tr.wg.Add(1)
				childctx, _ := context.WithCancel(ctx)
				go func(ctx context.Context) {
					defer tr.wg.Done()
					// Panic recovery wrapper
					defer func() {
						if r := recover(); r != nil {
							tr.Log.Error("Panic occurred in service", r, castedValue.GetName())
						}
					}()
					// blocking the current goroutine
					castedValue.Run(ctx)
				}(childctx)

			default:
				tr.Log.Info("value is of an unknown type: %T\n", castedValue)

			}
		}
	}

}

func (tr TaskRunner) GetName() string {
	return tr.Name
}
