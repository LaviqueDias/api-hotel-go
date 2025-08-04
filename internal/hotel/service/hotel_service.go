package service

import "github.com/LaviqueDias/api-hotel-go/internal/hotel/repository"

func NewHotelServiceInterface(repository repository.HotelRepositoryInterface) HotelServiceInterface {
	return &hotelServiceInterface{
		repository: repository,
	}
}

type hotelServiceInterface struct {
	repository repository.HotelRepositoryInterface
}

type HotelServiceInterface interface {

}