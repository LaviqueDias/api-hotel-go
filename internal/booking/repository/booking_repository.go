package repository

import "database/sql"

func NewBookingRepositoryInterface(databaseConnection *sql.DB) BookingRepositoryInterface {
	return &bookingRepositoryInterface{
		databaseConnection: databaseConnection,
	}
}

type bookingRepositoryInterface struct {
	databaseConnection *sql.DB
}

type BookingRepositoryInterface interface {

}