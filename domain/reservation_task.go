package domain

import "time"

type RESERVATION_TASK struct {
	ID             int       `json:"id" db:"id"`
	RESERVATION_ID int       `json:"-" db:"reservation"`
	RESERVATION    int       `json:"reservation" db:"-"`
	STAFF_ID       int       `json:"staffId" db:"staff_id"`
	STAFF          int       `json:"staff" db:"-"`
	STATUS         bool      `json:"status" db:"status"`
	DATE           time.Time `json:"date" db:"date"`
	UPDATED_AT     time.Time `json:"updatedAt" db:"updated_at"`
	CREATED_AT     time.Time `json:"createdAt" db:"created_at"`
}

type ReservationTaskRepository interface {
	Create(task *RESERVATION_TASK) error
	GetByReservationID(reservationID int) (*[]RESERVATION_TASK, error)
	Update(task *RESERVATION_TASK) error
}

type ReservationTaskUsecase interface {
	Create(task *RESERVATION_TASK) error
	GetByReservationID(reservationID int) (*[]RESERVATION_TASK, error)
	Update(task *RESERVATION_TASK) error
}
