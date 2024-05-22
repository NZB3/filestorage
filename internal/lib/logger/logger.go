package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
}

type logger struct {
	*logrus.Logger
}

func NewLogger() Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{})
	log.SetOutput(os.Stdout)

	return &logger{
		log,
	}
}
