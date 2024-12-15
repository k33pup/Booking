package logger

import (
	"github.com/k33pup/Booking/payment_svc/internal/pkg/config"
	"log/slog"
	"os"
)

func NewLogger() (*slog.Logger, error) {
	logFilePath, err := config.LoadLogFilePath()
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	logH := slog.NewJSONHandler(logFile, &slog.HandlerOptions{})
	return slog.New(logH), nil
}
