package dependecies

import (
	"database/sql"

	"github.com/LaviqueDias/api-hotel-go/internal/room/controller"
	"github.com/LaviqueDias/api-hotel-go/internal/room/repository"
	hotelRepository "github.com/LaviqueDias/api-hotel-go/internal/hotel/repository"
	"github.com/LaviqueDias/api-hotel-go/internal/room/service"
)

func InitRoomDependencies(database *sql.DB) controller.RoomControllerInterface {
	repository := repository.NewRoomRepositoryInterface(database)
	hotelRepository := hotelRepository.NewHotelReposirotyInterface(database)
	service := service.NewRoomServiceInterface(repository, hotelRepository)
	return controller.NewRoomControllerInterface(service)
}