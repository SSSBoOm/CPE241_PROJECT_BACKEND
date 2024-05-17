package payload

import "time"

type GetRoomAvailableGroupByRoomTypeDTO struct {
	START_DATE time.Time `json:"startDate" validate:"required"`
	END_DATE   time.Time `json:"endDate" validate:"required"`
}
