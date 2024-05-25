package domain

import "time"

type RoomType struct {
	ID          int       `json:"id" db:"id"`
	NAME        string    `json:"name" db:"name"`
	DETAIL      string    `json:"detail" db:"detail"`
	PRICE       *float32  `json:"price" db:"price"`
	ROOM        *[]ROOM   `json:"room,omitempty" db:"-"`
	ACCOMMODATE int       `json:"accommodate" db:"accommodate"`
	SIZE        string    `json:"size" db:"size"`
	BED         string    `json:"bed" db:"bed"`
	IMAGE_URL   string    `json:"imageUrl" db:"imageURL"`
	IS_ACTIVE   bool      `json:"isActive" db:"is_active"`
	UPDATED_AT  time.Time `json:"updateAt" db:"updated_at"`
	CREATED_AT  time.Time `json:"createdAt" db:"created_at"`
}

type RoomTypeRepository interface {
	GetAll() (*[]RoomType, error)
	GetByID(id int) (*RoomType, error)
	Create(roomType *RoomType) (*int, error)
	Update(roomType *RoomType) error
	UpdateIsActive(id int, isActive bool) error
}

type RoomTypeUsecase interface {
	GetAll() (*[]RoomType, error)
	GetByID(id int) (*RoomType, error)
	Create(roomType *RoomType) (*int, error)
	Update(roomType *RoomType) error
	UpdateIsActive(id int, isActive bool) error
}
