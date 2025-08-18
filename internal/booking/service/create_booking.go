package service

import (
	"github.com/LaviqueDias/api-hotel-go/internal/booking/model"
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/logger"
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/rest_err"
	"go.uber.org/zap"
)

func (bs *bookingServiceInterface) CreateBooking(bookingDTO *model.BookingDTO, roomIDs []int) (*model.BookingDTO, *rest_err.RestErr) {
	logger.Info("Init CreateBooking service",
		zap.String("journey", "CreateBooking"),
	)

	bookingDTO, err := bs.repository.CreateBooking(bookingDTO, roomIDs)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "CreateBooking"))
		return nil, err
	}

	logger.Info("CreateBooking service executed succesfully",
		zap.String("journey", "CreateBooking"),
	)

	return bookingDTO, err

}