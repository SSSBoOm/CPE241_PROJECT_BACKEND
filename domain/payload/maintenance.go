package payload

import "github.com/SSSBoOm/CPE241_Project_Backend/domain"

type MaintenanceCreateDTO struct {
	ROOM_ID         int                                      `json:"roomId" validate:"required"`
	MAINTENANCE_LOG *[]Maintenance_Log_On_Maintenance_Create `json:"maintenanceLog" validate:"omitempty"`
}
type Maintenance_Log_On_Maintenance_Create struct {
	DESCRIPTION string                        `json:"description" validate:"required"`
	STATUS      domain.MAINTENANCE_LOG_STATUS `json:"status" validate:"required"`
}
