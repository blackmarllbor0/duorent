package logger

import (
	run_mode "duorent.ru/pkg/run-mode"
	"github.com/sirupsen/logrus"
)

type LoggerService interface {
	Info(args ...interface{})
	Error(args ...interface{})
	Warn(args ...interface{})
}

func getLevelByAppRunMode(runMode string) uint32 {
	switch runMode {
	case run_mode.DEV, run_mode.TEST:
		return uint32(logrus.DebugLevel)
	case run_mode.PROD:
		return uint32(logrus.InfoLevel)
	default:
		return uint32(logrus.DebugLevel)
	}
}
