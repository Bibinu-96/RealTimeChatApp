package logger

import (
	"sync"

	"github.com/sirupsen/logrus"
)

var loggerInstance *LogrusLogger
var once sync.Once

type LogrusLogger struct {
	logger *logrus.Logger
}

func GetLogrusLogger() *LogrusLogger {
	once.Do(func() {
		loggerInstance = &LogrusLogger{logger: logrus.New()}
	})

	return loggerInstance
}

func (l *LogrusLogger) Info(msg string, args ...interface{}) {
	l.logger.Infof(msg, args...)
}

func (l *LogrusLogger) Warn(msg string, args ...interface{}) {
	l.logger.Warnf(msg, args...)
}

func (l *LogrusLogger) Error(msg string, args ...interface{}) {
	l.logger.Errorf(msg, args...)
}

func (l *LogrusLogger) Debug(msg string, args ...interface{}) {
	l.logger.Debugf(msg, args...)
}

func (l *LogrusLogger) Fatal(msg string, args ...interface{}) {
	l.logger.Fatalf(msg, args...)
}
