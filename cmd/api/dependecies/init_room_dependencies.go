package dependecies

import (
	"database/sql"

	"github.com/LaviqueDias/api-hotel-go/internal/room/controller"
	"github.com/LaviqueDias/api-hotel-go/internal/room/repository"
	"github.com/LaviqueDias/api-hotel-go/internal/room/service"
)

func InitRoomDependencies(database *sql.DB) controller.RoomControllerInterface {
	repository := repository.NewRoomRepositoryInterface(database)
	service := service.NewRoomServiceInterface(repository)
	return controller.NewRoomControllerInterface(service)
}