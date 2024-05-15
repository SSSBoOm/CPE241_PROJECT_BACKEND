package payload

import "github.com/SSSBoOm/CPE241_Project_Backend/domain"

type MAINTENANCE_CREATE struct {
	ROOM_ID         int                                      `json:"room_id" db:"room_id" validate:"required"`
	MAINTENANCE_LOG *[]MAINTENANCE_LOG_ON_MAINTENANCE_CREATE `json:"maintenance_log" validate:"omitempty"`
}

type MAINTENANCE_LOG_ON_MAINTENANCE_CREATE struct {
	DESCRIPTION string                        `json:"description" db:"description" validate:"required"`
	STATUS      domain.MAINTENANCE_LOG_STATUS `json:"status" db:"status" validate:"required"`
}
