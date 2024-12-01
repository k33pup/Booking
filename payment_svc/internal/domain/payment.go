package domain

type Payment struct {
	Price  float64 `json:"price"`
	RoomId string  `json:"room_id"`
}
