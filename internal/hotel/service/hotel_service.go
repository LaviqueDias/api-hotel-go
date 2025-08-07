package service

import (
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/rest_err"
	"github.com/LaviqueDias/api-hotel-go/internal/hotel/model"
	"github.com/LaviqueDias/api-hotel-go/internal/hotel/repository"
)

func NewHotelServiceInterface(repository repository.HotelRepositoryInterface) HotelServiceInterface {
	return &hotelServiceInterface{
		repository: repository,
	}
}

type hotelServiceInterface struct {
	repository repository.HotelRepositoryInterface
}

type HotelServiceInterface interface {
	CreateHotel(hotelDTO *model.HotelDTO) (*model.HotelDTO, *rest_err.RestErr)
	GetHotelByID(id int) (*model.HotelDTO, *rest_err.RestErr)
}