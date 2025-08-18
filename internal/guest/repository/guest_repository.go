package repository

import (
	"database/sql"

	"github.com/LaviqueDias/api-hotel-go/internal/configuration/rest_err"
	"github.com/LaviqueDias/api-hotel-go/internal/guest/model"
)

func NewGuestRepositoryInterface(databaseConnection *sql.DB) GuestRepositoryInterface {
	return &guestRepositoryInterface{
		databaseConnection: databaseConnection,
	}
}

type guestRepositoryInterface struct {
	databaseConnection *sql.DB
}

type GuestRepositoryInterface interface {
	CreateGuest(guestDTO *model.GuestDTO) (*model.GuestDTO, *rest_err.RestErr)
	GetGuestByID(idGuest int) (*model.GuestDTO, *rest_err.RestErr)
}