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

func (tr *TaskRunner) Run(ctx context.Context) error {
	tr.Log.Info("Running TaskRunner")
	errChan := make(chan error, 3)
	statusChannel := make(chan string, 3)
	taskChannel := channels.GetTaskChannel()

	// Ensure proper cleanup
	defer close(errChan)
	defer close(statusChannel)

	for {
		select {
		case <-ctx.Done():
			tr.Log.Info("Context canceled", "taskRunner", tr.Name)
			tr.wg.Wait()
			return errors.New("context canceled")
		case err := <-errChan:
			tr.Log.Error("Error occurred", "error", err)
		case status := <-statusChannel:
			tr.Log.Info("Status update", "status", status)
		case config := <-taskChannel:
			switch castedValue := config.(type) {
			case Task:
				tr.Log.Info("Executing task")
				go func(errCh chan error, statusCh chan string) {
					defer func() {
						if r := recover(); r != nil {
							errCh <- fmt.Errorf("panic occurred: %v", r)
						}
					}()
					castedValue.Invoke(errCh, statusCh)
				}(errChan, statusChannel)
			case PeriodicTask:
				tr.Log.Info("Initializing periodic task", "task", castedValue.Action.Name)
				tr.wg.Add(1)
				childCtx, cancel := context.WithCancel(ctx)
				go func(ctx context.Context) {
					defer tr.wg.Done()
					defer cancel() // Ensure context cleanup
					defer func() {
						if r := recover(); r != nil {
							tr.Log.Error("Panic occurred in periodic task", "task", castedValue.GetName(), "panic", r)
						}
					}()
					err := castedValue.Run(ctx)
					if err != nil {
						errChan <- err
					}
				}(childCtx)
			default:
				tr.Log.Info("Received unknown type", "type", fmt.Sprintf("%T", castedValue))
			}
		}
	}
}

func (tr *TaskRunner) GetName() string {
	return tr.Name
}
