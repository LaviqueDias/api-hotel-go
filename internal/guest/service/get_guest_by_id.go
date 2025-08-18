package service

import (
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/logger"
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/rest_err"
	"github.com/LaviqueDias/api-hotel-go/internal/guest/model"
	"go.uber.org/zap"
)

func (gs *guestServiceInterface) GetGuestByID(idGuest int) (*model.GuestDTO, *rest_err.RestErr) {
	logger.Info("Init GetGuestByID service",
		zap.Int("idGuest", idGuest),
		zap.String("journey", "GetGuestByID"),
	)

	guestDTO, err := gs.repository.GetGuestByID(idGuest)
	if err != nil {
		logger.Error("Error trying to call repository",
			err, zap.Int("idGuest", idGuest), zap.String("journey", "GetGuestByID"))
		return nil, err
	}

	logger.Info("GetGuestByID service executed successfully",
		zap.String("journey", "GetGuestByID"),
	)

	return guestDTO, nil
}