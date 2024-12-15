package models

import "github.com/k33pup/Booking/payment_svc/internal/domain"

func ToDomainPayment(payment PaymentM) domain.Payment {
	return domain.Payment{
		Amount:     payment.Amount,
		RoomId:     payment.RoomId,
		WebhookUrl: payment.WebhookUrl,
	}
}
