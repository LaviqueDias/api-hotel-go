package controller

import (
	"net/http"

	"github.com/LaviqueDias/api-hotel-go/internal/booking/model"
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/logger"
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (bc *bookingControllerInterface) CreateBooking(c *gin.Context) {
	logger.Info("Init CreateBooking controller",
		zap.String("journey", "CreateBooking"),
	)
	var BookingRequest model.BookingRequest

	if err := c.ShouldBindJSON(&BookingRequest); err != nil {
		logger.Error("Error trying to validate Booking info", err,
			zap.String("journey", "CreateBooking"),
		)
		errRest := rest_err.NewBadRequestError("Incorrect params")

		c.JSON(errRest.Code, errRest)
		return
	}

	BookingDTO := model.BookingRequestToBookingDTO(BookingRequest)

	BookingCreatedDTO, err := bc.service.CreateBooking(BookingDTO, BookingRequest.RoomIDs)
	if err != nil {
		logger.Error("Error trying to call CreateBooking service", err,
			zap.String("journey", "CreateBooking"),
		)
		c.JSON(err.Code, err)
	}

	BookingResponse := model.BookingDTOToBookingResponse(BookingCreatedDTO)

	logger.Info("CreateBooking controller executed succesfully",
		zap.String("journey", "CreateBooking"),
	)

	c.JSON(http.StatusOK, BookingResponse)

}