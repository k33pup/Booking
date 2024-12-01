package domain

type Result struct {
	IsPaid bool   `json:"result"`
	RoomId string `json:"room_id"`
}
