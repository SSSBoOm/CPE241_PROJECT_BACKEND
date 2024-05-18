package domain

import (
	"database/sql"
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

type RESERVATION struct {
	ID              int                `json:"id" db:"id"`
	ROOM_ID         int                `json:"room_id" db:"room_id"`
	USER_ID         string             `json:"user_id" db:"user_id"`
	START_DATE      time.Time          `json:"start_date" db:"start_date"`
	END_DATE        time.Time          `json:"end_date" db:"end_date"`
	PRICE           float64            `json:"price" db:"price"`
	STATUS          RESERVATION_STATUS `json:"status" db:"status"`
	PAYMENT_DATE    time.Time          `json:"payment_date" db:"payment_date"`
	PAYMENT_INFO_ID int                `json:"payment_info_id" db:"payment_info_id"`
	PAYMENT_INFO    Payment            `json:"payment_info" db:"-"`
	STAFF_ID        *sql.NullString    `json:"staff_id" db:"staff_id"`
	CREATED_AT      time.Time          `json:"created_at" db:"created_at"`
	UPDATED_AT      time.Time          `json:"updated_at" db:"updated_at"`
}

type ReservationUsecase interface {
	GetAll() (*[]RESERVATION, error)
	GetByDate(startDate string, endDate string) (*[]RESERVATION, error)
	GetByUserID(userID int) (*[]RESERVATION, error)
	GetByRoomID(roomID int) (*[]RESERVATION, error)
	GetByID(id int) (*RESERVATION, error)
	Create(reservation *RESERVATION) (id *int, err error)
	Update(reservation *RESERVATION) error
}

type ReservationRepository interface {
	GetAll() (*[]RESERVATION, error)
	GetByDate(startDate string, endDate string) (*[]RESERVATION, error)
	GetByUserID(userID int) (*[]RESERVATION, error)
	GetByRoomID(roomID int) (*[]RESERVATION, error)
	GetByID(id int) (*RESERVATION, error)
	Create(reservation *RESERVATION) (id *int, err error)
	Update(reservation *RESERVATION) error
}
