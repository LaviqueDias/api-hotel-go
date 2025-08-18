package service

import (
	"github.com/LaviqueDias/api-hotel-go/internal/booking/model"
	"github.com/LaviqueDias/api-hotel-go/internal/booking/repository"
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/rest_err"
)

func NewBookingServiceInterface(repository repository.BookingRepositoryInterface) BookingServiceInterface {
	return &bookingServiceInterface{
		repository: repository,
	}
}

type bookingServiceInterface struct {
	repository repository.BookingRepositoryInterface
}

type BookingServiceInterface interface {
	CreateBooking(bookingDTO *model.BookingDTO, roomIDs []int) (*model.BookingDTO, *rest_err.RestErr)
}