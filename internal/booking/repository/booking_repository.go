package repository

import (
	"database/sql"

	"github.com/LaviqueDias/api-hotel-go/internal/booking/model"
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/rest_err"
)

func NewBookingRepositoryInterface(databaseConnection *sql.DB) BookingRepositoryInterface {
	return &bookingRepositoryInterface{
		databaseConnection: databaseConnection,
	}
}

type bookingRepositoryInterface struct {
	databaseConnection *sql.DB
}

type BookingRepositoryInterface interface {
	CreateBooking(bookingDTO *model.BookingDTO, roomIDs []int) (*model.BookingDTO, *rest_err.RestErr)
}