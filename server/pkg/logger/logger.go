package logger

//go:generate mockgen -destination=./mock/logger.go -package=mock . Logger
type Logger interface {
	Printf(format string, v ...any)
	Println(v ...any)
}
