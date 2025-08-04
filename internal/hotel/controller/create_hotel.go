package controller

import (
	"net/http"

	"github.com/LaviqueDias/api-hotel-go/internal/configuration/logger"
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/rest_err"
	"github.com/LaviqueDias/api-hotel-go/internal/hotel/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (hc *hotelControllerInterface) CreateHotel(c *gin.Context) {
	logger.Info("Init CreateHotel controller",
		zap.String("journey", "CreateHotel"),
	)
	var hotelRequest model.HotelRequest

	if err := c.ShouldBindJSON(&hotelRequest); err != nil{
		logger.Error("Error trying to validate hotel info", err,
		zap.String("journey", "CreateHotel"),
	)
		errRest := rest_err.NewBadRequestError("Incorrect params")

		c.JSON(errRest.Code, errRest)
		return
	}

	hotelDTO := model.FromHotelRequest(hotelRequest)

	hotelCreatedDTO, err := hc.service.CreateHotel(hotelDTO)
	if err != nil{
		logger.Error("Error trying to call CreateHotel service", err,
		zap.String("journey", "CreateHotel"),
	)
		c.JSON(err.Code, err)
	}

	hotelResponse := model.ToHotelResponse(hotelCreatedDTO)

	logger.Info("CreateHotel controller executed succesfully", 
		zap.String("journey", "CreateHotel"),
	)

	c.JSON(http.StatusOK, hotelResponse)

}