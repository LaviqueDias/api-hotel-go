package model

type RoomDTO struct  {
	ID         int    
	RoomNumber int    
	Status     string 
	HotelID    int    
}

func RoomToRoomDTO(room *Room) *RoomDTO {
	return &RoomDTO{
		ID:         room.ID,
		RoomNumber: room.RoomNumber,
		Status:     room.Status,
		HotelID:    room.HotelID,
	}
}

func RoomDTOToRoom(roomDTO *RoomDTO) *Room {
	return &Room{
		ID:         roomDTO.ID,
		RoomNumber: roomDTO.RoomNumber,
		Status:     roomDTO.Status,
		HotelID:    roomDTO.HotelID,
	}
}

func RoomRequestToRoomDTO(roomRequest RoomRequest) *RoomDTO {
	return &RoomDTO{
		RoomNumber: roomRequest.RoomNumber,
		Status: roomRequest.Status,
		HotelID: roomRequest.HotelID,
	}
}

func  RoomDTOToRoomResponse(roomDTO *RoomDTO) RoomResponse {
	return RoomResponse{
		ID: roomDTO.ID,
		RoomNumber: roomDTO.RoomNumber,
		Status: roomDTO.Status,
		HotelID: roomDTO.HotelID,
	}
}