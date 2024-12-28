package logger

import "github.com/sirupsen/logrus"

type LogrusLogger struct {
	logger *logrus.Logger
}

func NewLogrusLogger() *LogrusLogger {
	return &LogrusLogger{logger: logrus.New()}
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
