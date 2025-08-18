package routes

import (
	"github.com/LaviqueDias/api-hotel-go/internal/guest/controller"
	"github.com/gin-gonic/gin"
)

func InitGuestRoutes(r *gin.RouterGroup, guestController controller.GuestControllerInterface) {

	r.POST("/", guestController.CreateGuest)
	// r.GET("/list", guestController.Listguests)
	r.GET("/id/:id", guestController.GetGuestByID)
	// r.PUT("/:id", guestController.Updateguest)
	// r.DELETE("/:id", guestController.Deleteguest)

}