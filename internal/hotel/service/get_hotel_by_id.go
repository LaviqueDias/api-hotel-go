package service

import (
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/logger"
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/rest_err"
	"github.com/LaviqueDias/api-hotel-go/internal/hotel/model"
	"go.uber.org/zap"
)

func (hs *hotelServiceInterface) GetHotelByID(id int) (*model.HotelDTO, *rest_err.RestErr) {
	logger.Info("Init GetHotelByID service",
		zap.Int("id", id),
		zap.String("journey", "GetHotelByID"),
	)

	hotelDTO, err := hs.repository.GetHotelByID(id)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.Int("id", id),
			zap.String("journey", "GetHotelByID"))
		return nil, err
	}

	logger.Info("GetHotelByID service executed successfully",
		zap.Int("id", id),
		zap.String("journey", "GetHotelByID"),
	)

	return hotelDTO, nil
}