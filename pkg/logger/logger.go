package logger

import (
	"log/slog"
)

//mock logger -> Grafana (future)

type Logger struct {
}

func (*Logger) LogError(str string) {
	slog.Error(str)
}

func (*Logger) LogInfo(str string) {
	slog.Info(str)
}
