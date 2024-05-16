package payload

type RoomCreateDTO struct {
	ROOM_NUMBER  string `json:"room_number" validate:"required"`
	IS_ACTIVE    bool   `json:"is_active" validate:"required"`
	ROOM_TYPE_ID int    `json:"room_type_id" validate:"required"`
}

type RoomUpdateDTO struct {
	ID           int    `json:"id" validate:"required"`
	ROOM_NUMBER  string `json:"room_number" validate:"required"`
	IS_ACTIVE    bool   `json:"is_active" validate:"required"`
	ROOM_TYPE_ID int    `json:"room_type_id" validate:"required"`
}
