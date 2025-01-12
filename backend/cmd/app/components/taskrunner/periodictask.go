package taskrunner

import "time"

type PeriodicTask struct {
	Action Task
	Ticker *time.Ticker
}
