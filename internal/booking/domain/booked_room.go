package domain

import "time"

type BookedRoom struct {
	ID          string
	HotelID     string
	Name        string
	Description string
	Price       int
	Entry       time.Time
	Exit        time.Time
	Email       string
}
