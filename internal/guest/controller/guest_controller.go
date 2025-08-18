package controller

import (
	"github.com/LaviqueDias/api-hotel-go/internal/guest/service"
	"github.com/gin-gonic/gin"
)

func NewGuestControllerInterface(service service.GuestServiceInterface) GuestControllerInterface {
	return &guestControllerInterface{
		service: service,
	}
}

type guestControllerInterface struct {
	service service.GuestServiceInterface
}

type GuestControllerInterface interface {
	CreateGuest(c *gin.Context)
	GetGuestByID(c *gin.Context)
}