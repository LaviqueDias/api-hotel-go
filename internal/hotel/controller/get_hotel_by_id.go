package controller

import (
	"net/http"
	"strconv"

	"github.com/LaviqueDias/api-hotel-go/internal/configuration/logger"
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/rest_err"
	"github.com/LaviqueDias/api-hotel-go/internal/hotel/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (hc *hotelControllerInterface) GetHotelByID(c *gin.Context) {
	logger.Info("Init GetHotelByID controller",
		zap.String("journey", "GetHotelByID"),
	)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		restErr := rest_err.NewBadRequestError("invalid id hotel parameter")
		logger.Error("Invalid idHotel parameter", restErr,
			zap.String("journey", "GetHotelByID"),
		)
		c.JSON(restErr.Code, restErr)
		return
	}

	logger.Info("Calling service GetHotelByID",
		zap.Int("hotelId", id),
		zap.String("journey", "GetHotelByID"),
	)

	hotelDTO, restErr := hc.service.GetHotelByID(id)
	if restErr != nil {
		logger.Error("Error returned from service GetHotelByID", restErr,
			zap.Int("hotelId", id),
			zap.String("journey", "GetHotelByID"),
		)
		c.JSON(restErr.Code, restErr)
		return
	}

	logger.Info("Hotel successfully retrieved",
		zap.Int("hotelId", hotelDTO.ID),
		zap.String("hotelName", hotelDTO.Name),
		zap.String("journey", "GetHotelByID"),
	)

	hotelResponse := model.HotelDTOToHotelResponse(hotelDTO)

	c.JSON(http.StatusOK, hotelResponse)
}
