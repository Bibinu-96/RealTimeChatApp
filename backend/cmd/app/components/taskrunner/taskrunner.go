package taskrunner

import (
	"backend/internal/channels"
	"backend/pkg/logger"
	"context"
	"errors"
	"fmt"
)

type TaskRunner struct {
	Log  logger.Logger
	Name string
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
			default:
				tr.Log.Info("value is of an unknown type: %T\n", castedValue)

			}
		}
	}

}

func (tr TaskRunner) GetName() string {
	return tr.Name
}
