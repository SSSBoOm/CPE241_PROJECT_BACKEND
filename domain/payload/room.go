package payload

type RoomCreateDTO struct {
	ROOM_NUMBER  string `json:"roomNumber" validate:"required"`
	IS_ACTIVE    bool   `json:"isActive" validate:"boolean"`
	ROOM_TYPE_ID int    `json:"roomTypeId" validate:"required"`
}

type RoomUpdateDTO struct {
	ID           int    `json:"id" validate:"required"`
	ROOM_NUMBER  string `json:"roomNumber" validate:"required"`
	IS_ACTIVE    bool   `json:"isActive" validate:"boolean"`
	ROOM_TYPE_ID int    `json:"roomTypeId" validate:"required"`
}

type RoomUpdateRoomIsActiveDTO struct {
	ID       int  `json:"id" validate:"required"`
	IsActive bool `json:"isActive" validate:"boolean"`
}
