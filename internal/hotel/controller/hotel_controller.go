package controller

import "github.com/LaviqueDias/api-hotel-go/internal/hotel/service"

func NewHoteControllerInterface(service service.HotelServiceInterface) HotelControllerInterface {
	return &hotelControllerInterface{
		service: service,
	}
}

type hotelControllerInterface struct {
	service service.HotelServiceInterface
}

type HotelControllerInterface interface {
}