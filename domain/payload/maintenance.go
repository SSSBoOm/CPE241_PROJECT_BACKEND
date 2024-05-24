package payload

import (
	"time"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
)

type MaintenanceCreateDTO struct {
	TITLE           string                                   `json:"title" validate:"required"`
	ROOM_ID         int                                      `json:"roomId" validate:"required"`
	MAINTENANCE_LOG *[]Maintenance_Log_On_Maintenance_Create `json:"maintenanceLog" validate:"omitempty"`
}

type Maintenance_Log_On_Maintenance_Create struct {
	DESCRIPTION string                        `json:"description" validate:"required"`
	DATE        time.Time                     `json:"date" validate:"required"`
	STATUS      domain.MAINTENANCE_LOG_STATUS `json:"status" validate:"required"`
}
