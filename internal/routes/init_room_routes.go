package routes

import (
	"github.com/LaviqueDias/api-hotel-go/internal/room/controller"
	"github.com/gin-gonic/gin"
)

func InitRoomRoutes(r *gin.RouterGroup, roomController controller.RoomControllerInterface) {
	r.POST("/", roomController.CreateRoom)
}