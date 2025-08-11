package service

import (
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/logger"
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/rest_err"
	"github.com/LaviqueDias/api-hotel-go/internal/hotel/model"
	roomModel "github.com/LaviqueDias/api-hotel-go/internal/room/model"
	"go.uber.org/zap"
)

func (hs *hotelServiceInterface) GetHotelByID(id int) (*model.HotelDTO, *rest_err.RestErr) {
	logger.Info("Init GetHotelByID service",
		zap.Int("id", id),
		zap.String("journey", "GetHotelByID"),
	)

	hotelDTO, err := hs.repository.GetHotelByID(id)
	if err != nil {
		logger.Error("Error trying to call repository",
			err, zap.Int("id", id), zap.String("journey", "GetHotelByID"))
		return nil, err
	}

	// Busca os quartos desse hotel
	rooms, rerr := hs.roomRepository.GetRoomsByHotelID(id)
	if rerr != nil {
		logger.Error("Error trying to fetch rooms for hotel",
			rerr, zap.Int("hotelId", id), zap.String("journey", "GetHotelByID"))
		return nil, rest_err.NewInternalServerError("error fetching hotel rooms")
	}

	// Garanta slice não-nil (JSON sai [] ao invés de null)
	if rooms == nil {
		rooms = []roomModel.Room{}
	}
	hotelDTO.Rooms = rooms

	logger.Info("GetHotelByID service executed successfully",
		zap.Int("id", id), zap.Int("roomsCount", len(rooms)),
		zap.String("journey", "GetHotelByID"),
	)

	return hotelDTO, nil
}