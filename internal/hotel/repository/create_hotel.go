package repository

import (
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/logger"
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/rest_err"
	"github.com/LaviqueDias/api-hotel-go/internal/hotel/model"
	"go.uber.org/zap"
)

func (hr *hotelRepositoryInterface) CreateHotel(hotelDTO *model.HotelDTO) (*model.HotelDTO, *rest_err.RestErr) {
	logger.Info("Init CreateHotel repository",
		zap.String("journey", "CreateHotel"),
	)

	hotel := model.HotelDTOToHotel(hotelDTO)

	query := `INSERT INTO hotels (name) 
              VALUES ($1)
              RETURNING id`

	var lastInsertID int
	err := hr.databaseConnection.QueryRow(query, hotel.Name).Scan(&lastInsertID)
	if err != nil {
		logger.Error("Error trying to insert hotel",
			err,
			zap.String("journey", "CreateHotel"),
		)
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	hotel.ID = lastInsertID

	hotelDTO = model.HotelToHotelDTO(hotel)

	logger.Info("Hotel inserted successfully",
		zap.Int("hotelId", lastInsertID),
		zap.String("journey", "CreateHotel"),
	)

	return hotelDTO, nil
}
