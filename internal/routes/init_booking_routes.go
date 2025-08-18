package routes

import (
	"github.com/LaviqueDias/api-hotel-go/internal/booking/controller"
	"github.com/gin-gonic/gin"
)

func InitBookingRoutes(r *gin.RouterGroup, bookingController controller.BookingControllerInterface) {
	r.POST("/", bookingController.CreateBooking)
}