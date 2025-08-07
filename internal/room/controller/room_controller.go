package controller

import (
	"github.com/LaviqueDias/api-hotel-go/internal/room/service"
	"github.com/gin-gonic/gin"
)

func NewRoomControllerInterface(service service.RoomServiceInterface) RoomControllerInterface {
	return &roomControllerInterface{
		service: service,
	}
}

type roomControllerInterface struct {
	service service.RoomServiceInterface
}

type RoomControllerInterface interface {
	CreateRoom(c *gin.Context)
}