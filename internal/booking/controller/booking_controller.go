package controller

import (
	"github.com/LaviqueDias/api-hotel-go/internal/booking/service"
	"github.com/gin-gonic/gin"
)

func NewBookingControllerInterface(service service.BookingServiceInterface) BookingControllerInterface {
	return &bookingControllerInterface{
		service: service,
	}
}

type bookingControllerInterface struct {
	service service.BookingServiceInterface
}

type BookingControllerInterface interface {
	CreateBooking(c *gin.Context)
}