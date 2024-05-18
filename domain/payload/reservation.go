package payload

import "time"

type CreateReservationDTO struct {
	ROOM_ID         int       `json:"roomId" validate:"required"`
	START_DATE      time.Time `json:"startDate" validate:"required"`
	END_DATE        time.Time `json:"endDate" validate:"required"`
	PRICE           float64   `json:"price" validate:"required"`
	PAYMENT_INFO_ID int       `json:"paymentInfoId" validate:"required"`
}

type GetRoomAvailableGroupByRoomTypeDTO struct {
	START_DATE time.Time `json:"startDate" validate:"required"`
	END_DATE   time.Time `json:"endDate" validate:"required"`
}
