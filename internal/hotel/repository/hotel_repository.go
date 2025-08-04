package repository

import (
	"database/sql"

)

func NewHotelReposirotyInterface(databaseConnection *sql.DB) HotelRepositoryInterface {
	return &hotelRepositoryInterface{
		databaseConnection: databaseConnection,
	}
}

type hotelRepositoryInterface struct {
	databaseConnection *sql.DB
}

type HotelRepositoryInterface interface {

}