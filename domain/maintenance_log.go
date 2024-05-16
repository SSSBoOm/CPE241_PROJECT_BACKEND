package domain

type MAINTENANCE_LOG_STATUS string

const (
	CASE_OPEN MAINTENANCE_LOG_STATUS = "CASE_OPEN"
	PENDING   MAINTENANCE_LOG_STATUS = "PENDING"
	APPROVED  MAINTENANCE_LOG_STATUS = "APPROVED"
	REJECTED  MAINTENANCE_LOG_STATUS = "REJECTED"
)

type MAINTENANCE_LOG struct {
	ID             int                    `json:"id" db:"id"`
	MAINTENANCE_ID int                    `json:"maintenance_id" db:"maintenance_id"`
	STAFF_ID       string                 `json:"staff_id" db:"staff_id"`
	DESCRIPTION    string                 `json:"description" db:"description"`
	STATUS         MAINTENANCE_LOG_STATUS `json:"status" db:"status"`
	CREATED_AT     string                 `json:"created_at" db:"created_at"`
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
