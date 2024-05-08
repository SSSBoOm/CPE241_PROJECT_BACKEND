package payload

type CreateRoomType struct {
	NAME      string `json:"name" db:"name" validate:"required"`
	IS_ACTIVE bool   `json:"isActive" db:"is_active" validate:"required"`
}

type UpdateRoomType struct {
	ID        int    `json:"id" db:"id" validate:"required"`
	NAME      string `json:"name" db:"name" validate:"required"`
	IS_ACTIVE bool   `json:"isActive" db:"is_active" validate:"required"`
}
