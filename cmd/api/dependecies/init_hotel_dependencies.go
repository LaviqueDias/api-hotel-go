package dependecies

import (
	"database/sql"

	"github.com/LaviqueDias/api-hotel-go/internal/hotel/controller"
	"github.com/LaviqueDias/api-hotel-go/internal/hotel/repository"
	roomRepository "github.com/LaviqueDias/api-hotel-go/internal/room/repository"
	"github.com/LaviqueDias/api-hotel-go/internal/hotel/service"
)

func InitHotelDependencies(database *sql.DB) controller.HotelControllerInterface {
	repository := repository.NewHotelReposirotyInterface(database)
	roomRepository := roomRepository.NewRoomRepositoryInterface(database)
	service := service.NewHotelServiceInterface(repository, roomRepository)
	return controller.NewHotelControllerInterface(service)

}