package repository

import (
	"github.com/LaviqueDias/api-hotel-go/internal/booking/model"
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/logger"
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/rest_err"
	"go.uber.org/zap"
)

func (br *bookingRepositoryInterface) CreateBooking(bookingDTO *model.BookingDTO, roomIDs []int) (*model.BookingDTO, *rest_err.RestErr) {
	logger.Info("Init CreateBooking repository",
		zap.String("journey", "CreateBooking"),
	)

	tx, err := br.databaseConnection.Begin()
	if err != nil {
		logger.Error("failed to begin tx", err, zap.String("journey", "CreateBooking"))
		return nil, rest_err.NewInternalServerError("could not start transaction")
	}
	defer func() { _ = tx.Rollback() }()

	const insertBooking = `
		INSERT INTO bookings (checkin_date, checkout_date, status, total_price, guest_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`

	var bookingID int
	if err := tx.QueryRow(
		insertBooking,
		bookingDTO.CheckinDate,
		bookingDTO.CheckoutDate,
		bookingDTO.Status,
		bookingDTO.TotalPrice,
		bookingDTO.GuestID,
	).Scan(&bookingID); err != nil {
		logger.Error("failed to insert booking", err, zap.String("journey", "CreateBooking"))
		return nil, rest_err.NewInternalServerError("error inserting booking")
	}

	// 2) INSERTs em booking_rooms
	if len(roomIDs) > 0 {
		const insertBookingRoom = `
			INSERT INTO booking_room (booking_id, room_id)
			VALUES ($1, $2)
			ON CONFLICT DO NOTHING; -- opcional, se tiver PK(booking_id, room_id)
		`

		stmt, err := tx.Prepare(insertBookingRoom)
		if err != nil {
			logger.Error("failed to prepare insert booking_rooms", err, zap.String("journey", "CreateBooking"))
			return nil, rest_err.NewInternalServerError("error preparing booking_rooms insertion")
		}
		defer stmt.Close()

		for _, roomID := range roomIDs {
			if _, err := stmt.Exec(bookingID, roomID); err != nil {
				logger.Error("failed to insert into booking_rooms", err,
					zap.Int("room_id", roomID),
					zap.Int("booking_id", bookingID),
					zap.String("journey", "CreateBooking"),
				)
				return nil, rest_err.NewInternalServerError("error linking room to booking")
			}
		}
	}

	// 3) commit
	if err := tx.Commit(); err != nil {
		logger.Error("failed to commit tx", err, zap.String("journey", "CreateBooking"))
		return nil, rest_err.NewInternalServerError("could not commit transaction")
	}

	bookingDTO.ID = bookingID
	logger.Info("CreateBooking repository executed successfully",
		zap.Int("booking_id", bookingDTO.ID),
		zap.Int("rooms_linked", len(roomIDs)),
		zap.String("journey", "CreateBooking"),
	)
	return bookingDTO, nil
}
