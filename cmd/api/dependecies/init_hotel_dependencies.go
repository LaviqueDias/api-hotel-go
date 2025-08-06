package dependecies

import (
	"database/sql"

	"github.com/LaviqueDias/api-hotel-go/internal/hotel/controller"
	"github.com/LaviqueDias/api-hotel-go/internal/hotel/repository"
	"github.com/LaviqueDias/api-hotel-go/internal/hotel/service"
)

func InitHotelDependencies(database *sql.DB) controller.HotelControllerInterface {
	repository := repository.NewHotelReposirotyInterface(database)
	service := service.NewHotelServiceInterface(repository)
	return controller.NewHoteControllerInterface(service)

}