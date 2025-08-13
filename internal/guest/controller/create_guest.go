package controller

import (
	"net/http"

	"github.com/LaviqueDias/api-hotel-go/internal/configuration/logger"
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/rest_err"
	"github.com/LaviqueDias/api-hotel-go/internal/guest/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (hc *guestControllerInterface) CreateGuest(c *gin.Context) {
	logger.Info("Init CreateGuest controller",
		zap.String("journey", "CreateGuest"),
	)
	var guestRequest model.GuestRequest

	if err := c.ShouldBindJSON(&guestRequest); err != nil {
		logger.Error("Error trying to validate Guest info", err,
			zap.String("journey", "CreateGuest"),
		)
		errRest := rest_err.NewBadRequestError("Incorrect params")

		c.JSON(errRest.Code, errRest)
		return
	}

	guestDTO := model.GuestRequestToGuestDTO(guestRequest)

	guestCreatedDTO, err := hc.service.CreateGuest(guestDTO)
	if err != nil {
		logger.Error("Error trying to call CreateGuest service", err,
			zap.String("journey", "CreateGuest"),
		)
		c.JSON(err.Code, err)
	}

	GuestResponse := model.GuestDTOToGuestResponse(guestCreatedDTO)

	logger.Info("CreateGuest controller executed succesfully",
		zap.String("journey", "CreateGuest"),
	)

	c.JSON(http.StatusOK, GuestResponse)

}