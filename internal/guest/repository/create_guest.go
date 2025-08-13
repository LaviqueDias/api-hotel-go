package repository

import (
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/logger"
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/rest_err"
	"github.com/LaviqueDias/api-hotel-go/internal/guest/model"
	"go.uber.org/zap"
)

func (gr *guestRepositoryInterface) CreateGuest(guestDTO *model.GuestDTO) (*model.GuestDTO, *rest_err.RestErr) {
	logger.Info("Init CreateGuest repository",
		zap.String("journey", "CreateGuest"),
	)

	guest := model.GuestDTOToGuest(guestDTO)

	query := `INSERT INTO guests (name) 
              VALUES ($1)
              RETURNING id`

	var lastInsertID int
	err := gr.databaseConnection.QueryRow(query, guest.Name).Scan(&lastInsertID)
	if err != nil {
		logger.Error("Error trying to insert guest",
			err,
			zap.String("journey", "CreateGuest"),
		)
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	guest.ID = lastInsertID

	guestDTO = model.GuestToGuestDTO(guest)

	logger.Info("guest inserted successfully",
		zap.Int("guestId", lastInsertID),
		zap.String("journey", "CreateGuest"),
	)

	return guestDTO, nil
}