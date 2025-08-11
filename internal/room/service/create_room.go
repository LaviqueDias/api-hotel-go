package service

import (
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/logger"
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/rest_err"
	"github.com/LaviqueDias/api-hotel-go/internal/room/model"
	"go.uber.org/zap"
)

func (rs *roomServiceInterface) CreateRoom(roomDTO *model.RoomDTO) (*model.RoomDTO, *rest_err.RestErr) {
	logger.Info("Init CreateRoom service",
		zap.String("journey", "CreateRoom"),
	)

	// valida hotel existente antes de inserir
	if _, err := rs.hotelRepository.GetHotelByID(roomDTO.HotelID); err != nil {
		logger.Error("Hotel not found when creating room", err,
			zap.Int("hotelId", roomDTO.HotelID), zap.String("journey", "CreateRoom"))
		return nil, rest_err.NewNotFoundError("hotel not found")
	}

	createdRoomDTO, err := rs.repository.CreateRoom(roomDTO)
	if err != nil {
		logger.Error("Error trying to create room", err,
			zap.String("journey", "CreateRoom"))
		return nil, err
	}

	logger.Info("CreateRoom service executed successfully",
		zap.Int("roomId", createdRoomDTO.ID),
		zap.String("journey", "CreateRoom"),
	)

	return createdRoomDTO, nil
}