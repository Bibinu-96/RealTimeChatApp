package channels

var taskChannel chan interface{}

type Job func(chan error, chan string)

func SetTaskChannel(jobChannel chan interface{}) {
	taskChannel = jobChannel
}

func GetTaskChannel() chan interface{} {
	return taskChannel
}
