package service

import (
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/rest_err"
	"github.com/LaviqueDias/api-hotel-go/internal/room/model"
	"github.com/LaviqueDias/api-hotel-go/internal/room/repository"
	hotelRepository "github.com/LaviqueDias/api-hotel-go/internal/hotel/repository"
)

func NewRoomServiceInterface(repository repository.RoomRepositoryInterface, hotelRepository hotelRepository.HotelRepositoryInterface) RoomServiceInterface {
	return &roomServiceInterface{
		repository: repository,
		hotelRepository: hotelRepository,
	}
}

type roomServiceInterface struct {
	repository repository.RoomRepositoryInterface
	hotelRepository hotelRepository.HotelRepositoryInterface
}

type RoomServiceInterface interface {
	CreateRoom(roomDTO *model.RoomDTO) (*model.RoomDTO, *rest_err.RestErr)
}