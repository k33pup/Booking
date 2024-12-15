package models

type PaymentM struct {
	Amount     float64 `json:"amount"`
	RoomId     string  `json:"room_id"`
	WebhookUrl string  `json:"webhook_url"`
}
