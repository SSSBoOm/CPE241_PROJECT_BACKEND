package domain

import "time"

type RoomType struct {
	ID        int       `json:"id" db:"id"`
	NAME      string    `json:"name" db:"name"`
	IS_ACTIVE bool      `json:"isActive" db:"is_active"`
	UPDATE_AT time.Time `json:"updateAt" db:"update_at"`
	CREATE_AT time.Time `json:"createdAt" db:"created_at"`
}

type RoomTypeRepository interface {
	GetAll() (*[]RoomType, error)
	GetByID(id int) (*RoomType, error)
	Create(roomType *RoomType) error
	Update(roomType *RoomType) error
}

type RoomTypeUsecase interface {
	GetAll() (*[]RoomType, error)
	GetByID(id int) (*RoomType, error)
	Create(roomType *RoomType) error
	Update(roomType *RoomType) error
}
