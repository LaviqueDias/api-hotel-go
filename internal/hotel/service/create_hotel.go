package service

import (
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/logger"
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/rest_err"
	"github.com/LaviqueDias/api-hotel-go/internal/hotel/model"
	"go.uber.org/zap"
)

func (hs *hotelServiceInterface) CreateHotel(hotelDTO *model.HotelDTO) (*model.HotelDTO, *rest_err.RestErr){
	logger.Info("Init CreateHotel service",
		zap.String("journey", "CreateHotel"),
	)

	hotelDTO, err := hs.repository.CreateHotel(hotelDTO)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "CreateHotel"))
		return nil, err
	}
	
	logger.Info("CreateHotel service executed succesfully", 
		zap.String("journey", "CreateHotel"),
	)
	
	return hotelDTO, err

}