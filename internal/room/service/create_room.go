package service

import (
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/logger"
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/rest_err"
	"github.com/LaviqueDias/api-hotel-go/internal/room/model"
	"go.uber.org/zap"
)

func (rs *roomServiceInterface) CreateRoom(roomDTO *model.RoomDTO) (*model.RoomDTO, *rest_err.RestErr) {
	logger.Info("Init Createroom service",
		zap.String("journey", "Createroom"),
	)
	
	roomDTO, err := rs.repository.CreateRoom(roomDTO)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "CreateRoom"))
		return nil, err
	}

	logger.Info("CreateRoom service executed succesfully",
		zap.String("journey", "CreateRoom"),
	)

	return roomDTO, err

}