package service

import (
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/rest_err"
	"github.com/LaviqueDias/api-hotel-go/internal/room/model"
	"github.com/LaviqueDias/api-hotel-go/internal/room/repository"
)

func NewRoomServiceInterface(repository repository.RoomRepositoryInterface) RoomServiceInterface {
	return &roomServiceInterface{
		repository: repository,
	}
}

type roomServiceInterface struct {
	repository repository.RoomRepositoryInterface
}

type RoomServiceInterface interface {
	CreateRoom(roomDTO *model.RoomDTO) (*model.RoomDTO, *rest_err.RestErr)
}