package payload

import "time"

type ReservationTaskCreateDTO struct {
	RESERVATION_ID int       `json:"reservationId" validate:"required"`
	DATE           time.Time `json:"date" validate:"required"`
}

type ReservationTaskUpdateStaffDTO struct {
	ID int `json:"id" validate:"required"`
}

type ReservationTaskUpdateStatusDTO struct {
	ID     int  `json:"id" validate:"required"`
	STATUS bool `json:"status" validate:"boolean"`
}
