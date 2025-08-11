package repository

import (
	"database/sql"

	"github.com/LaviqueDias/api-hotel-go/internal/configuration/rest_err"
	"github.com/LaviqueDias/api-hotel-go/internal/room/model"
)

func NewRoomRepositoryInterface(databaseConnection *sql.DB) RoomRepositoryInterface {
	return &roomRepositoryInterface{
		databaseConnection: databaseConnection,
	}
}

type roomRepositoryInterface struct {
	databaseConnection *sql.DB
}

type RoomRepositoryInterface interface {
	CreateRoom(roomDTO *model.RoomDTO) (*model.RoomDTO, *rest_err.RestErr)
	GetRoomsByHotelID(hotelID int) ([]model.Room, *rest_err.RestErr)
}