package domain

type MAINTENANCE struct {
	ID              int                `json:"id" db:"id"`
	ROOM_ID         int                `json:"room_id" db:"room_id"`
	STAFF_ID        int                `json:"staff_id" db:"staff_id"`
	MAINTENANCE_LOG *[]MAINTENANCE_LOG `json:"maintenance_log" db:"-"`
	CREATED_AT      string             `json:"created_at" db:"created_at"`
}

type MaintenanceRepository interface {
	Get(id int) (*MAINTENANCE, error)
	GetAll() (*[]MAINTENANCE, error)
	Create(maintenance *MAINTENANCE) error
	Update(maintenance *MAINTENANCE) error
}

type MaintenanceUsecase interface {
	Get(id int) (*MAINTENANCE, error)
	GetAll() (*[]MAINTENANCE, error)
	Create(maintenance *MAINTENANCE) error
	Update(maintenance *MAINTENANCE) error
}
