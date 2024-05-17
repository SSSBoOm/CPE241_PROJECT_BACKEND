package payload

type CreateRoomType struct {
	NAME      string `json:"name" validate:"required"`
	IS_ACTIVE bool   `json:"isActive" validate:"required"`
}

type UpdateRoomType struct {
	ID        int    `json:"id" validate:"required"`
	NAME      string `json:"name" validate:"required"`
	IS_ACTIVE bool   `json:"isActive" validate:"required"`
}
