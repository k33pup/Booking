package openapi

import (
	"github.com/k33pup/Booking/payment_svc/internal/domain"
	clientGen "github.com/k33pup/Booking/payment_svc/internal/pkg/transport/http/generated_api/generated_client"
)

func ToDomainPayment(request CreatePaymentPostRequest) *domain.Payment {
	return &domain.Payment{
		RoomId:     request.RoomId,
		Amount:     float64(request.Amount),
		WebhookUrl: request.WebhookUrl,
	}
}

func FromDomainPayment(model domain.Payment) *clientGen.CreatePaymentPostRequest {
	amount := float32(model.Amount)
	return &clientGen.CreatePaymentPostRequest{
		RoomId:     &model.RoomId,
		Amount:     &amount,
		WebhookUrl: &model.WebhookUrl,
	}
}

type ResultResponse struct {
	RoomId string `json:"room_id"`
	Status string `json:"status"`
}

func ToResultResponse(result *domain.Payment) *ResultResponse {
	return &ResultResponse{RoomId: result.RoomId, Status: "proceeding"}
}
