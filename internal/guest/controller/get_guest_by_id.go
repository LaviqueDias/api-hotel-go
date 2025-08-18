package controller

import (
	"net/http"
	"strconv"

	"github.com/LaviqueDias/api-hotel-go/internal/configuration/logger"
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/rest_err"
	"github.com/LaviqueDias/api-hotel-go/internal/guest/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (gc *guestControllerInterface) GetGuestByID(c *gin.Context) {
	logger.Info("Init GetGuestByID controller",
		zap.String("journey", "GetGuestByID"),
	)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		rest_err := rest_err.NewBadRequestError("invalid id guest parameter")
		logger.Error("Invalid idGuest parameter", rest_err,
			zap.String("journey", "GetGuestByID"),
		)
		c.JSON(rest_err.Code, rest_err)
		return 
	}

	logger.Info("Calling service GetGuestByID",
		zap.Int("hotelId", id),
		zap.String("journey", "GetGuestByID"),
	)

	guestDTO, restErr := gc.service.GetGuestByID(id)
	if restErr != nil {
		logger.Error("Error returned from service GetGuestByID", restErr,
			zap.Int("guestId", id),
			zap.String("journey", "GuestByID"),
		)
		c.JSON(restErr.Code, restErr)
		return		
	}

	logger.Info("Guest successfully retrieved",
		zap.Int("guestId", guestDTO.ID),
		zap.String("GuestName", guestDTO.Name),
		zap.String("journey", "GetGuestByID"),
	)

	guestResponse := model.GuestDTOToGuestResponse(guestDTO)

	c.JSON(http.StatusOK, guestResponse)

}