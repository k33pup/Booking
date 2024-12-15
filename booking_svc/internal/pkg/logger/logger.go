package logger

import (
	"github.com/k33pup/Booking.git/internal/pkg/config"
	"log/slog"
	"os"
)

func NewLogger() (*slog.Logger, error) {
	logFilePath, err := config.LoadLogFilePath()
	if err != nil {
		return nil, err
	}
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	handler := slog.NewTextHandler(logFile, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	log := slog.New(handler)
	return log, nil
}
