package dependecies

import (
	"database/sql"

	"github.com/LaviqueDias/api-hotel-go/internal/booking/controller"
	"github.com/LaviqueDias/api-hotel-go/internal/booking/repository"
	"github.com/LaviqueDias/api-hotel-go/internal/booking/service"
)

func InitBookingDependencies(database *sql.DB) controller.BookingControllerInterface {
	repository := repository.NewBookingRepositoryInterface(database)
	service := service.NewBookingServiceInterface(repository)
	return controller.NewBookingControllerInterface(service)

}