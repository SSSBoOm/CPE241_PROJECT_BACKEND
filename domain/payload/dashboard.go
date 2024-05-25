package payload

import "time"

type GetDashboardReservationDTO struct {
	START_DATE time.Time `json:"startDate" validate:"required"`
	END_DATE   time.Time `json:"endDate" validate:"required"`
}
