package postgres

import "github.com/k33pup/Booking/hotel_svc/internal/domain"

func ToDomainHotel(model *HotelModel) *domain.Hotel {
	return &domain.Hotel{
		Id:          model.Id,
		Name:        model.Name,
		Description: model.Description,
		City:        model.City,
		Address:     model.Address,
		Contacts:    model.Contacts,
	}
}

func ToListDomainHotel(models []*HotelModel) []*domain.Hotel {
	var hotels []*domain.Hotel
	for _, hotel := range models {
		addHotel := ToDomainHotel(hotel)
		hotels = append(hotels, addHotel)
	}
	return hotels
}

func ToRepositoryHotel(domain *domain.Hotel) *HotelModel {
	return &HotelModel{
		Id:          domain.Id,
		Name:        domain.Name,
		Description: domain.Description,
		City:        domain.City,
		Address:     domain.Address,
		Contacts:    domain.Contacts,
	}
}

func ToDomainRoom(model *RoomModel) *domain.Room {
	return &domain.Room{
		Id:          model.Id,
		HotelId:     model.HotelId,
		Name:        model.Name,
		Description: model.Description,
		Price:       model.Price,
	}
}

func ToListDomainRoom(models []*RoomModel) []*domain.Room {
	var rooms []*domain.Room
	for _, room := range models {
		addRoom := ToDomainRoom(room)
		rooms = append(rooms, addRoom)
	}
	return rooms
}

func ToRepositoryRoom(domain *domain.Room) *RoomModel {
	return &RoomModel{
		Id:          domain.Id,
		HotelId:     domain.HotelId,
		Name:        domain.Name,
		Description: domain.Description,
		Price:       domain.Price,
	}
}
