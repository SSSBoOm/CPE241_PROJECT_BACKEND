package domain

import "time"

type RESERVATION_TASK struct {
	ID             int       `json:"id" db:"id"`
	RESERVATION_ID int       `json:"-" db:"reservation_id"`
	RESERVATION    int       `json:"reservation" db:"-"`
	STAFF_ID       *string   `json:"staffId" db:"staff_id"`
	STAFF          *User     `json:"staff" db:"-"`
	STATUS         bool      `json:"status" db:"status"`
	DATE           time.Time `json:"date" db:"date"`
	UPDATED_AT     time.Time `json:"updatedAt" db:"updated_at"`
	CREATED_AT     time.Time `json:"createdAt" db:"created_at"`
}

type ReservationTaskRepository interface {
	Create(task *RESERVATION_TASK) error
	GetByReservationID(reservationID int) (*RESERVATION_TASK, error)
	Update(task *RESERVATION_TASK) error
	UpdateStaff(id int, staffID string) error
	UpdateStatus(id int, status bool) error
}

type ReservationTaskUsecase interface {
	Create(task *RESERVATION_TASK) error
	GetByReservationID(reservationID int) (*RESERVATION_TASK, error)
	Update(task *RESERVATION_TASK) error
	UpdateStaff(id int, staffID string) error
	UpdateStatus(id int, status bool) error
}
