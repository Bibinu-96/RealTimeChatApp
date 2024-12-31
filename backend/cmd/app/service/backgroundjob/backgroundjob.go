package backgroundjob

import (
	"backend/internal/channels"
	"backend/pkg/logger"
	"context"
	"errors"
	"fmt"
)

type BackgroundJob struct {
	Log  logger.Logger
	Name string
}

func (bg BackgroundJob) Run(ctx context.Context) error {
	bg.Log.Info("Running background Jobs")
	errChan := make(chan error, 3)
	statusChannel := make(chan string, 3)
	taskChannel := channels.GetTaskChannel()
	for {

		select {
		case <-ctx.Done():
			bg.Log.Info("context cancelled", bg.Name)
			return errors.New("context cancelled")
		case err := <-errChan:
			bg.Log.Error("err occured", err)
			//return err
		case status := <-statusChannel:
			bg.Log.Info("status", status)

		case config := <-taskChannel:
			switch castedValue := config.(type) {
			case channels.Job:
				bg.Log.Info(" executing the task")
				go func(errCh chan error, statusCh chan string) {
					defer func() {
						if r := recover(); r != nil {
							// Handle the panic and send the error to the error channel
							errCh <- fmt.Errorf("panic occurred: %v", r)
						}
					}()
					castedValue(errCh, statusCh)
				}(errChan, statusChannel)
			default:
				bg.Log.Info("value is of an unknown type: %T\n", castedValue)

			}
		}
	}

}

func (bg BackgroundJob) GetName() string {
	return bg.Name
}
