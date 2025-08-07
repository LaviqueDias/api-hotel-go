package repository

import (
	"database/sql"
	"fmt"

	"github.com/LaviqueDias/api-hotel-go/internal/configuration/logger"
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/rest_err"
	"github.com/LaviqueDias/api-hotel-go/internal/hotel/model"
	"go.uber.org/zap"
)

func (hr *hotelRepositoryInterface) GetHotelByID(id int) (*model.HotelDTO, *rest_err.RestErr) {
	logger.Info("Init GetHotelByID repository",
		zap.Int("hotelId", id),
		zap.String("journey", "GetHotelByID"),
	)

	query := `SELECT id, name FROM hotels WHERE id = $1`
	var hotel model.Hotel

	err := hr.databaseConnection.QueryRow(query, id).Scan(&hotel.ID, &hotel.Name)
	if err != nil {
        if err == sql.ErrNoRows {
            logger.Info("Hotel not found by id",
                zap.String("journey", "GetHotelByID"),
            )
            return nil, rest_err.NewNotFoundError(fmt.Sprintf("hotel with id %d not found", id))
        }
        logger.Error("Error scanning hotel by id",
            err,
            zap.String("journey", "GetHotelByID"),
        )
        return nil, rest_err.NewInternalServerError("error finding hotel by id")
    }

	hotelDTO := model.HotelToHotelDTO(&hotel)

	logger.Info("Hotel retrieved successfully",
		zap.Int("hotelId", hotel.ID),
		zap.String("journey", "GetHotelByID"),
	)

	return hotelDTO, nil
}