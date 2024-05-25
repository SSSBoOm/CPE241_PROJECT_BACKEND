package payload

import (
	"time"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
)

type MaintenanceLogCreate struct {
	MAINTENANCE_ID int                           `json:"maintenanceId" validate:"required"`
	DESCRIPTION    string                        `json:"description" validate:"required"`
	DATE           time.Time                     `json:"date" validate:"required"`
	IMAGE_URL      string                        `json:"imageUrl"`
	STATUS         domain.MAINTENANCE_LOG_STATUS `json:"status" validate:"required"`
}
