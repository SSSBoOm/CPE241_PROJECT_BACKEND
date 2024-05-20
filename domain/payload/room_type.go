package payload

type CreateRoomType struct {
	NAME      string  `json:"name" validate:"required"`
	DETAIL    string  `json:"detail" validate:"required"`
	PRICE     float32 `json:"price" validate:"required"`
	IS_ACTIVE bool    `json:"isActive" validate:"boolean"`
}

type UpdateRoomType struct {
	ID        int     `json:"id" validate:"required"`
	NAME      string  `json:"name" validate:"required"`
	DETAIL    string  `json:"detail" validate:"required"`
	PRICE     float32 `json:"price" validate:"required"`
	IS_ACTIVE bool    `json:"isActive" validate:"boolean"`
}

type UpdateRoomTypeIsActiveDTO struct {
	ID       int  `json:"id" validate:"required"`
	IS_ACTIVE bool `json:"isActive" validate:"boolean"`
}
