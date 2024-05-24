package domain

import "time"

type MAINTENANCE_LOG_STATUS string

const (
	MAINTENANCE_LOG_STATUS_CASE_OPEN MAINTENANCE_LOG_STATUS = "CASE_OPEN"
	MAINTENANCE_LOG_STATUS_PENDING   MAINTENANCE_LOG_STATUS = "PENDING"
	MAINTENANCE_LOG_STATUS_APPROVED  MAINTENANCE_LOG_STATUS = "APPROVED"
	MAINTENANCE_LOG_STATUS_REJECTED  MAINTENANCE_LOG_STATUS = "REJECTED"
	MAINTENANCE_LOG_STATUS_DONE      MAINTENANCE_LOG_STATUS = "DONE"
)

type MAINTENANCE_LOG struct {
	ID             int                    `json:"id" db:"id"`
	MAINTENANCE_ID int                    `json:"-" db:"maintenance_id"`
	STAFF_ID       string                 `json:"staffId" db:"staff_id"`
	STAFF          *User                  `json:"staff" db:"-"`
	DESCRIPTION    string                 `json:"description" db:"description"`
	DATE           time.Time              `json:"date" db:"date"`
	STATUS         MAINTENANCE_LOG_STATUS `json:"status" db:"status"`
	UPDATED_AT     time.Time              `json:"updatedAt" db:"updated_at"`
	CREATED_AT     time.Time              `json:"createdAt" db:"created_at"`
}

type MaintenanceLogRepository interface {
	GetByID(id int) (*MAINTENANCE_LOG, error)
	GetAll() (*[]MAINTENANCE_LOG, error)
	GetByMaintenanceID(maintenance_id int) (*[]MAINTENANCE_LOG, error)
	Create(maintenanceLog *MAINTENANCE_LOG) error
	Update(maintenanceLog *MAINTENANCE_LOG) error
}

type MaintenanceLogUsecase interface {
	GetByID(id int) (*MAINTENANCE_LOG, error)
	GetAll() (*[]MAINTENANCE_LOG, error)
	GetByMaintenanceID(maintenance_id int) (*[]MAINTENANCE_LOG, error)
	Create(maintenanceLog *MAINTENANCE_LOG) error
	Update(maintenanceLog *MAINTENANCE_LOG) error
}
