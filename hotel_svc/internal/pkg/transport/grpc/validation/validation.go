package validation

import (
	"errors"
	"github.com/go-playground/validator/v10"
	apiv1pb "github.com/k33pup/Booking/hotel_svc/internal/pkg/transport/grpc/generated"
)

var validate = validator.New()

func ValidateField(field interface{}, tag string) error {
	return validate.Var(field, tag)
}

func ValidateCreateHotel(in *apiv1pb.CreateHotelRequest) error {
	err := ValidateField(in.Name, "required,min=1,max=255")
	if err != nil {
		return errors.New("validation error: " + err.Error())
	}
	if err = ValidateField(in.Description, "required,min=1,max=255"); err != nil {
		return errors.New("validation error: " + err.Error())
	}
	if err = ValidateField(in.City, "required,min=1,max=255"); err != nil {
		return errors.New("validation error: " + err.Error())
	}
	if err = ValidateField(in.Address, "required,min=1,max=255"); err != nil {
		return errors.New("validation error: " + err.Error())
	}
	if err = ValidateField(in.Contacts, "required,min=1,max=255"); err != nil {
		return errors.New("validation error: " + err.Error())
	}
	return nil
}

func ValidateCreateRoom(in *apiv1pb.CreateRoomRequest) error {
	err := ValidateField(in.HotelId, "required,min=1")
	if err != nil {
		return errors.New("validation error: " + err.Error())
	}
	if err = ValidateField(in.Name, "required,min=1,max=255"); err != nil {
		return errors.New("validation error: " + err.Error())
	}
	if err = ValidateField(in.Description, "required,min=1,max=255"); err != nil {
		return errors.New("validation error: " + err.Error())
	}
	if err = ValidateField(in.Price, "required,gt=0"); err != nil {
		return errors.New("validation error: " + err.Error())
	}
	return nil
}
