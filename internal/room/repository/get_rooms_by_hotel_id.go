package repository

import (
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/rest_err"
	"github.com/LaviqueDias/api-hotel-go/internal/room/model"
)

func (rr *roomRepositoryInterface) GetRoomsByHotelID(hotelID int) ([]model.Room, *rest_err.RestErr) {
	const q = `
		SELECT id, number, status, daily_price, hotel_id
		FROM rooms
		WHERE hotel_id = $1
		ORDER BY number
	`
	rows, err := rr.databaseConnection.Query(q, hotelID)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	var rooms []model.Room
	for rows.Next() {
		var r model.Room
		if err := rows.Scan(&r.ID, &r.RoomNumber, &r.Status, &r.DailyPrice, &r.HotelID); err != nil {
			return nil, rest_err.NewInternalServerError(err.Error())
		}
		rooms = append(rooms, r)
	}
	if err := rows.Err(); err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	return rooms, nil
}
