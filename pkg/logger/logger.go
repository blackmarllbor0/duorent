package logger

type LoggerService interface {
	Info(args ...interface{})
	Error(args ...interface{})
	Warn(args ...interface{})
}
