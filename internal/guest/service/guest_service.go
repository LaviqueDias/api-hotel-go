package service

import (
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/rest_err"
	"github.com/LaviqueDias/api-hotel-go/internal/guest/model"
	"github.com/LaviqueDias/api-hotel-go/internal/guest/repository"
)

func NewGuestServiceInterface(repository repository.GuestRepositoryInterface) GuestServiceInterface {
	return &guestServiceInterface{
		repository:     repository,
	}
}

type guestServiceInterface struct {
	repository     repository.GuestRepositoryInterface
}

type GuestServiceInterface interface {
	CreateGuest(GuestDTO *model.GuestDTO) (*model.GuestDTO, *rest_err.RestErr)
} 