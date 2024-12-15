package usecases

import (
	"context"
	"github.com/k33pup/Booking/payment_svc/internal/domain"
)

type IPaymentUseCase interface {
	CreatePayment(ctx context.Context, payment *domain.Payment) (*domain.Payment, error)
}

type PaymentUseCase struct{}

func NewPaymentUseCase() *PaymentUseCase {
	return &PaymentUseCase{}
}

func (h *PaymentUseCase) CreatePayment(ctx context.Context, payment *domain.Payment) (*domain.Payment, error) {
	return payment, nil
}
