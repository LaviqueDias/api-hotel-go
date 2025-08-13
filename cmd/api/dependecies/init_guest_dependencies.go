package dependecies

import (
	"database/sql"

	"github.com/LaviqueDias/api-hotel-go/internal/guest/controller"
	"github.com/LaviqueDias/api-hotel-go/internal/guest/repository"
	"github.com/LaviqueDias/api-hotel-go/internal/guest/service"
)

func InitGuestDependencies(database *sql.DB) controller.GuestControllerInterface {
	repository := repository.NewGuestRepositoryInterface(database)
	service := service.NewGuestServiceInterface(repository)
	return controller.NewGuestControllerInterface(service)

}