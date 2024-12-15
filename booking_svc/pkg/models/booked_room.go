package models

import "time"

type BookedRoom struct {
	ID      string
	HotelID string
	Entry   time.Time
	Exit    time.Time
	Email   string
	IsPaid  bool
}
