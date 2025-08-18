package repository

import (
	"database/sql"
	"fmt"

	"github.com/LaviqueDias/api-hotel-go/internal/configuration/logger"
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/rest_err"
	"github.com/LaviqueDias/api-hotel-go/internal/guest/model"
	"go.uber.org/zap"
)

func (gr *guestRepositoryInterface) GetGuestByID(idGuest int) (*model.GuestDTO, *rest_err.RestErr) {
	logger.Info("Init GetGuestByID repository",
		zap.Int("GuestId", idGuest),
		zap.String("journey", "GetGuestByID"),
	)

	query := `SELECT id, name FROM guests WHERE id = $1`
	var Guest model.Guest

	err := gr.databaseConnection.QueryRow(query, idGuest).Scan(&Guest.ID, &Guest.Name)
	if err != nil {
        if err == sql.ErrNoRows {
            logger.Info("Guest not found by id",
                zap.String("journey", "GetGuestByID"),
            )
            return nil, rest_err.NewNotFoundError(fmt.Sprintf("Guest with id %d not found", idGuest))
        }
        logger.Error("Error scanning Guest by id",
            err,
            zap.String("journey", "GetGuestByID"),
        )
        return nil, rest_err.NewInternalServerError("error finding Guest by id")
    }

	GuestDTO := model.GuestToGuestDTO(&Guest)

	logger.Info("Guest retrieved successfully",
		zap.Int("GuestId", Guest.ID),
		zap.String("journey", "GetGuestByID"),
	)

	return GuestDTO, nil
}