package domain

import (
	"time"
)

type RESERVATION_STATUS string

const (
	RESERVATION_STATUS_WAITING_APPROVE_PAYMENT RESERVATION_STATUS = "WAITING_APPROVE_PAYMENT"
	RESERVATION_STATUS_REJECTED_PAYMENT        RESERVATION_STATUS = "REJECTED_PAYMENT"
	RESERVATION_STATUS_APPROVED_PAYMENT        RESERVATION_STATUS = "APPROVED_PAYMENT"
	RESERVATION_STATUS_WAITING_CHECKIN         RESERVATION_STATUS = "WAITING_CHECKIN"
	RESERVATION_STATUS_CHECKED_IN              RESERVATION_STATUS = "CHECKED_IN"
	RESERVATION_STATUS_WAITING_CHECKED_OUT     RESERVATION_STATUS = "WAITING_CHECKED_OUT"
	RESERVATION_STATUS_SUCCESS                 RESERVATION_STATUS = "SUCCESS"
	RESERVATION_STATUS_CANCELED                RESERVATION_STATUS = "CANCELED"
)

type RESERVATION_TYPE string

const (
	RESERVATION_TYPE_ROOM    RESERVATION_TYPE = "ROOM"
	RESERVATION_TYPE_SERVICE RESERVATION_TYPE = "SERVICE"
)

type RESERVATION struct {
	ID              int                `json:"id" db:"id"`
	TYPE            RESERVATION_TYPE   `json:"type" db:"type"`
	ROOM_ID         *int               `json:"-" db:"room_id"`
	ROOM            *ROOM              `json:"room" db:"-"`
	SERVICE_ID      *int               `json:"-" db:"service_id"`
	SERVICE         *SERVICE           `json:"service" db:"-"`
	USER_ID         string             `json:"userId" db:"user_id"`
	START_DATE      time.Time          `json:"startDate" db:"start_date"`
	END_DATE        time.Time          `json:"endDate" db:"end_date"`
	PRICE           float64            `json:"price" db:"price"`
	STATUS          RESERVATION_STATUS `json:"status" db:"status"`
	PAYMENT_DATE    time.Time          `json:"paymentDate" db:"payment_date"`
	PAYMENT_INFO_ID int                `json:"-" db:"payment_info_id"`
	PAYMENT_INFO    *Payment           `json:"paymentInfo" db:"-"`
	STAFF_ID        *string            `json:"staffId" db:"staff_id"`
	STAFF           *User              `json:"staff" db:"-"`
	CREATED_AT      time.Time          `json:"createdAt" db:"created_at"`
	UPDATED_AT      time.Time          `json:"updatedAt" db:"updated_at"`
}

type ReservationUsecase interface {
	GetAll() (*[]RESERVATION, error)
	GetByDate(startDate string, endDate string) (*[]RESERVATION, error)
	GetByUserID(userID string) (*[]RESERVATION, error)
	GetByRoomID(roomID int) (*[]RESERVATION, error)
	GetByID(id int) (*RESERVATION, error)
	Create(reservation *RESERVATION) (id *int, err error)
	Update(reservation *RESERVATION) error
	UpdateStaff(id int, staffID string) error
	UpdateStatus(id int, status RESERVATION_STATUS) error
}

type ReservationRepository interface {
	GetAll() (*[]RESERVATION, error)
	GetByDate(startDate string, endDate string) (*[]RESERVATION, error)
	GetByUserID(userID string) (*[]RESERVATION, error)
	GetByRoomID(roomID int) (*[]RESERVATION, error)
	GetByID(id int) (*RESERVATION, error)
	Create(reservation *RESERVATION) (id *int, err error)
	Update(reservation *RESERVATION) error
	UpdateStaff(id int, staffID string) error
	UpdateStatus(id int, status RESERVATION_STATUS) error
}
