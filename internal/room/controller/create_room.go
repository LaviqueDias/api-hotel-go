package controller

import (
	"net/http"

	"github.com/LaviqueDias/api-hotel-go/internal/configuration/logger"
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/rest_err"
	"github.com/LaviqueDias/api-hotel-go/internal/room/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (rc *roomControllerInterface) CreateRoom(c *gin.Context) {
	logger.Info("Init CreateRoom controller",
		zap.String("journey", "CreateRoom"),
	)
	var roomRequest model.RoomRequest

	if err := c.ShouldBindJSON(&roomRequest); err != nil {
		logger.Error("Error trying to validate room info", err,
			zap.String("journey", "CreateRoom"),
		)
		errRest := rest_err.NewBadRequestError("Incorrect params")

		c.JSON(errRest.Code, errRest)
		return
	}

	roomDTO := model.RoomRequestToRoomDTO(roomRequest)

	roomCreatedDTO, err := rc.service.CreateRoom(roomDTO)
	if err != nil {
		logger.Error("Error trying to call CreateRoom service", err,
			zap.String("journey", "CreateRoom"),
		)
		c.JSON(err.Code, err)
	}

	roomResponse := model.RoomDTOToRoomResponse(roomCreatedDTO)

	logger.Info("CreateRoom controller executed succesfully",
		zap.String("journey", "CreateRoom"),
	)

	c.JSON(http.StatusOK, roomResponse)

}