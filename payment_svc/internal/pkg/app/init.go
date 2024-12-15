package app

import (
	"context"
	"github.com/k33pup/Booking/payment_svc/internal/pkg/logger"
	impl "github.com/k33pup/Booking/payment_svc/internal/pkg/transport/http"
	"github.com/k33pup/Booking/payment_svc/internal/usecases"
	"log/slog"
)

type PaymentService struct {
	server *impl.Server
	Log    *slog.Logger
}

func NewPaymentService() (*PaymentService, error) {
	log, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	useCase := usecases.NewPaymentUseCase()
	server := impl.NewServer(useCase)
	return &PaymentService{server: server, Log: log}, nil
}

func (b *PaymentService) Start(ctx context.Context) error {
	b.Log.Info("Payment service started")
	err := b.server.Start()
	if err != nil {
		return err
	}
	return nil
}

func (b *PaymentService) Stop(ctx context.Context) error {
	b.Log.Info("Payment service stopped")
	err := b.server.Stop(ctx)
	if err != nil {
		return err
	}
	return nil
}
