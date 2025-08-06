package routes

import (

	"github.com/LaviqueDias/api-hotel-go/internal/hotel/controller"
	"github.com/gin-gonic/gin"
)

func InitHotelRoutes(r *gin.RouterGroup, hotelController controller.HotelControllerInterface) {

	r.POST("/", hotelController.CreateHotel)
	// r.GET("/list", hotelController.ListHotels)
	// r.GET("/:id", hotelController.GetHotelByID)
	// r.PUT("/:id", hotelController.UpdateHotel)
	// r.DELETE("/:id", hotelController.DeleteHotel)

}
