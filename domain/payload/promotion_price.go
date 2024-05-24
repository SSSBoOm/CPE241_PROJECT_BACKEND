package payload

import "time"

type PromotionPriceCreateDTO struct {
	NAME         string    `json:"name" validate:"required"`
	PRICE        float64   `json:"price" validate:"required"`
	START_DATE   time.Time `json:"startDate" validate:"required"`
	END_DATE     time.Time `json:"endDate" validate:"required"`
	ROOM_TYPE_ID []int     `json:"roomTypeId" validate:"required"`
}
