package repository

import (
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/logger"
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/rest_err"
	"github.com/LaviqueDias/api-hotel-go/internal/room/model"
	"go.uber.org/zap"
)

func (rr *roomRepositoryInterface) CreateRoom(roomDTO *model.RoomDTO) (*model.RoomDTO, *rest_err.RestErr) {
	logger.Info("Init CreateRoom repository",
		zap.String("journey", "CreateRoom"),
	)

	query := `
		INSERT INTO rooms (room_number, status, hotel_id) 
		VALUES ($1, $2, $3)
		RETURNING id
	`

	var lastInsertID int
	err := rr.databaseConnection.QueryRow(
		query,
		roomDTO.RoomNumber,
		roomDTO.Status,
		roomDTO.HotelID,
	).Scan(&lastInsertID)

	if err != nil {
		logger.Error("Error trying to insert room",
			err,
			zap.String("journey", "CreateRoom"),
		)
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	roomDTO.ID = lastInsertID

	logger.Info("Room inserted successfully",
		zap.Int("roomId", lastInsertID),
		zap.String("journey", "CreateRoom"),
	)

	return roomDTO, nil
}
