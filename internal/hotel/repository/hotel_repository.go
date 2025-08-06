package repository

import (
	"database/sql"

	"github.com/LaviqueDias/api-hotel-go/internal/configuration/rest_err"
	"github.com/LaviqueDias/api-hotel-go/internal/hotel/model"
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
	CreateHotel(hotelDTO *model.HotelDTO) (*model.HotelDTO, *rest_err.RestErr)
}