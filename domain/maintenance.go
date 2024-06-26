package domain

import "time"

type MAINTENANCE struct {
	ID              int                `json:"id" db:"id"`
	TITLE           string             `json:"title" db:"title"`
	ROOM_ID         int                `json:"-" db:"room_id"`
	ROOM            *ROOM              `json:"room" db:"-"`
	STAFF_ID        string             `json:"staffId" db:"staff_id"`
	MAINTENANCE_LOG *[]MAINTENANCE_LOG `json:"maintenanceLog" db:"-"`
	UPDATED_AT      time.Time          `json:"updatedAt" db:"updated_at"`
	CREATED_AT      time.Time          `json:"createdAt" db:"created_at"`
}

type MaintenanceRepository interface {
	GetByID(id int) (*MAINTENANCE, error)
	GetAll() (*[]MAINTENANCE, error)
	Create(maintenance *MAINTENANCE) (*int, error)
	Update(maintenance *MAINTENANCE) error
}

type MaintenanceUsecase interface {
	GetByID(id int) (*MAINTENANCE, error)
	GetAll() (*[]MAINTENANCE, error)
	Create(maintenance *MAINTENANCE) error
	CreateWithMaintenance_Log(maintenance *MAINTENANCE) error
	Update(maintenance *MAINTENANCE) error
}
