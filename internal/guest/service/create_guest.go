package service

import (
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/logger"
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/rest_err"
	"github.com/LaviqueDias/api-hotel-go/internal/guest/model"
	"go.uber.org/zap"
)

func (hs *guestServiceInterface) CreateGuest(GuestDTO *model.GuestDTO) (*model.GuestDTO, *rest_err.RestErr) {
	logger.Info("Init CreateGuest service",
		zap.String("journey", "CreateGuest"),
	)

	GuestDTO, err := hs.repository.CreateGuest(GuestDTO)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "CreateGuest"))
		return nil, err
	}

	logger.Info("CreateGuest service executed succesfully",
		zap.String("journey", "CreateGuest"),
	)

	return GuestDTO, err

}