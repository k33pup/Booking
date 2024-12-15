package logger

import (
	"log/slog"
)

//mock logger -> Grafana (future)

type Logger struct {
}

func (*Logger) LogError(err error) {
	slog.Error(err.Error())
}

func (*Logger) LogInfo(str string) {
	slog.Info(str)
}
