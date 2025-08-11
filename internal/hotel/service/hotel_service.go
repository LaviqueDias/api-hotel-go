package service

import (
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/rest_err"
	"github.com/LaviqueDias/api-hotel-go/internal/hotel/model"
	"github.com/LaviqueDias/api-hotel-go/internal/hotel/repository"
	roomRepository "github.com/LaviqueDias/api-hotel-go/internal/room/repository"
)

func NewHotelServiceInterface(repository repository.HotelRepositoryInterface, roomRepository roomRepository.RoomRepositoryInterface) HotelServiceInterface {
	return &hotelServiceInterface{
		repository: repository,
		roomRepository: roomRepository,
	}
}

type hotelServiceInterface struct {
	repository repository.HotelRepositoryInterface
	roomRepository roomRepository.RoomRepositoryInterface
}

type HotelServiceInterface interface {
	CreateHotel(hotelDTO *model.HotelDTO) (*model.HotelDTO, *rest_err.RestErr)
	GetHotelByID(id int) (*model.HotelDTO, *rest_err.RestErr)
}