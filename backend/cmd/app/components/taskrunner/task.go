package taskrunner

type Task struct {
	Name   string
	Action func(chan error, chan string)
}

func (t Task) Invoke(errChan chan error, statusChan chan string) {
	t.Action(errChan, statusChan)
}

func (t Task) GetName() string {
	return t.Name
}
