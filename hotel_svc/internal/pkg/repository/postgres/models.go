package postgres

type HotelModel struct {
	Id          string `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	City        string `gorm:"not null"`
	Address     string `gorm:"not null"`
	Contacts    string `gorm:"not null"`
}

func (HotelModel) TableName() string {
	return "hotels"
}

type RoomModel struct {
	Id          string `gorm:"primaryKey"`
	HotelId     string `gorm:"not null"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Price       uint64 `gorm:"not null"`
}

func (RoomModel) TableName() string {
	return "rooms"
}
