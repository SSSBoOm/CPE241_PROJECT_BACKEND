package payload

type CreateRoomType struct {
	NAME        string                  `json:"name" validate:"required"`
	DETAIL      string                  `json:"detail" validate:"required"`
	ACCOMMODATE int                     `json:"accommodate" validate:"required"`
	PRICE       float32                 `json:"price" validate:"required"`
	IMAGE_URL   string                  `json:"imageUrl" validate:"required"`
	IS_ACTIVE   bool                    `json:"isActive" validate:"boolean"`
	ROOM        *[]RoomOnCreateRoomType `json:"room" validate:"omitempty"`
}

type RoomOnCreateRoomType struct {
	ROOM_NUMBER string `json:"roomNumber" validate:"required"`
	IS_ACTIVE   bool   `json:"isActive" validate:"boolean"`
}

type UpdateRoomType struct {
	ID          int     `json:"id" validate:"required"`
	NAME        string  `json:"name" validate:"required"`
	DETAIL      string  `json:"detail" validate:"required"`
	ACCOMMODATE int     `json:"accommodate" validate:"required"`
	PRICE       float32 `json:"price" validate:"required"`
	IMAGE_URL   string  `json:"imageUrl" validate:"required"`
	IS_ACTIVE   bool    `json:"isActive" validate:"boolean"`
}

type UpdateRoomTypeIsActiveDTO struct {
	ID        int  `json:"id" validate:"required"`
	IS_ACTIVE bool `json:"isActive" validate:"boolean"`
}
