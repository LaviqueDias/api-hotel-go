package controller

import (
	"github.com/LaviqueDias/api-hotel-go/internal/hotel/service"
	"github.com/gin-gonic/gin"
)

func NewHoteControllerInterface(service service.HotelServiceInterface) HotelControllerInterface {
	return &hotelControllerInterface{
		service: service,
	}
}

type hotelControllerInterface struct {
	service service.HotelServiceInterface
}

type HotelControllerInterface interface {
	CreateHotel(c *gin.Context)
	GetHotelByID(c *gin.Context)
}