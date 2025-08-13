package model

type GuestDTO struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
}

// Guest → GuestDTO
func GuestToGuestDTO(guest *Guest) *GuestDTO {
	return &GuestDTO{
		ID:   guest.ID,
		Name: guest.Name,
	}
}

// GuestDTO → Guest
func GuestDTOToGuest(guestDTO *GuestDTO) *Guest {
	return &Guest{
		ID:   guestDTO.ID,
		Name: guestDTO.Name,
	}
}

// GuestRequest → GuestDTO
func GuestRequestToGuestDTO(guestRequest GuestRequest) *GuestDTO {
	return &GuestDTO{
		Name: guestRequest.Name,
	}
}

// GuestDTO → GuestResponse
func GuestDTOToGuestResponse(guestDTO *GuestDTO) GuestResponse {
	return GuestResponse{
		ID:   guestDTO.ID,
		Name: guestDTO.Name,
	}
}
