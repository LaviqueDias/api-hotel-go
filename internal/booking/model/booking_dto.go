package model

import "time"

type BookingDTO struct {
	ID           int
	CheckinDate  time.Time
	CheckoutDate time.Time
	Status       string
	TotalPrice   float64
	GuestID      int
	RoomIDs      []int  
}

// Booking → BookingDTO
func BookingToBookingDTO(booking *Booking) *BookingDTO {
	return &BookingDTO{
		ID:           booking.ID,
		CheckinDate:  booking.CheckinDate,
		CheckoutDate: booking.CheckoutDate,
		Status:       booking.Status,
		TotalPrice:   booking.TotalPrice,
		GuestID:      booking.GuestID,
	}
}

// BookingDTO → Booking
func BookingDTOToBooking(bookingDTO *BookingDTO) *Booking {
	return &Booking{
		ID:           bookingDTO.ID,
		CheckinDate:  bookingDTO.CheckinDate,
		CheckoutDate: bookingDTO.CheckoutDate,
		Status:       bookingDTO.Status,
		TotalPrice:   bookingDTO.TotalPrice,
		GuestID:      bookingDTO.GuestID,
	}
}

// BookingRequest → BookingDTO
func BookingRequestToBookingDTO(bookingRequest BookingRequest) *BookingDTO {
	return &BookingDTO{
		CheckinDate:  bookingRequest.CheckinDate,
		CheckoutDate: bookingRequest.CheckoutDate,
		Status:       bookingRequest.Status,
		TotalPrice:   bookingRequest.TotalPrice,
		GuestID:      bookingRequest.GuestID,
	}
}

// BookingDTO → BookingResponse
func BookingDTOToBookingResponse(bookingDTO *BookingDTO) BookingResponse {
	return BookingResponse{
		ID:           bookingDTO.ID,
		CheckinDate:  bookingDTO.CheckinDate,
		CheckoutDate: bookingDTO.CheckoutDate,
		Status:       bookingDTO.Status,
		TotalPrice:   bookingDTO.TotalPrice,
		GuestID:      bookingDTO.GuestID,
	}
}
