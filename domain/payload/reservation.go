package payload

import (
	"time"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
)

type CreateReservationDTO struct {
	TYPE            domain.RESERVATION_TYPE `json:"type" db:"type"`
	ROOM_ID         *int                    `json:"roomId"`
	SERVICE_ID      *int                    `json:"serviceId"`
	START_DATE      time.Time               `json:"startDate" validate:"required"`
	END_DATE        time.Time               `json:"endDate" validate:"required"`
	PRICE           float64                 `json:"price" validate:"required"`
	PAYMENT_INFO_ID int                     `json:"paymentInfoId" validate:"required"`
}

type GetRoomAvailableGroupByRoomTypeDTO struct {
	START_DATE time.Time `json:"startDate" validate:"required"`
	END_DATE   time.Time `json:"endDate" validate:"required"`
}

type UpdateReservationStaffDTO struct {
	RESERVATION_ID int    `json:"reservationId" validate:"required"`
	STAFF_ID       string `json:"staffId" validate:"required"`
}

type UpdateReservationStatusDTO struct {
	RESERVATION_ID int                       `json:"reservationId" validate:"required"`
	STATUS         domain.RESERVATION_STATUS `json:"status" validate:"required"`
}

type UpdateReservationPaymentDTO struct {
	RESERVATION_ID  int `json:"reservationId" validate:"required"`
	PAYMENT_INFO_ID int `json:"paymentInfoId" validate:"required"`
}
