package domain

import "time"

type Room struct {
	ID           int       `json:"id" db:"id"`
	ROOM_NUMBER  string    `json:"room_number" db:"room_number"`
	IS_ACTIVE    bool      `json:"is_active" db:"is_active"`
	ROOM_TYPE_ID int       `json:"-" db:"room_type_id"`
	ROOM_TYPE    RoomType  `json:"room_type"`
	UPDATE_AT    time.Time `json:"updateAt" db:"update_at"`
	CREATE_AT    time.Time `json:"createdAt" db:"created_at"`
}

type RoomRepository interface {
	GetAll() (*[]Room, error)
	GetByID(id int) (*Room, error)
	Create(room *Room) error
	Update(room *Room) error
}

type RoomUsecase interface {
	GetAll() (*[]Room, error)
	GetByID(id int) (*Room, error)
	Create(room *Room) error
	Update(room *Room) error
}
