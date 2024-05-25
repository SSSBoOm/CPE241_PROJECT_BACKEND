package domain

import "time"

type DashboardReservation struct {
	ID    int    `json:"id" db:"id"`
	NAME  string `json:"name" db:"name"`
	TOTAL int    `json:"total" db:"total"`
}

type DashboardReservation2 struct {
	NAME  string `json:"name" db:"name"`
	TOTAL int    `json:"total" db:"total"`
}

type DashboardService struct {
}

type DashboardUsecase interface {
	GetDashboardRoomTypeReservation(startDate time.Time, endDate time.Time) (*[]DashboardReservation, error)
	GetDashboardServiceTypeReservation(startDate time.Time, endDate time.Time) (*[]DashboardReservation, error)
	GetDashboardReservationByPaymentType(startDate time.Time, endDate time.Time) (*[]DashboardReservation2, error)
}

type DashboardRepository interface {
	GetDashboardRoomTypeReservation(startDate time.Time, endDate time.Time) (*[]DashboardReservation, error)
	GetDashboardServiceTypeReservation(startDate time.Time, endDate time.Time) (*[]DashboardReservation, error)
	GetDashboardReservationByPaymentType(startDate time.Time, endDate time.Time) (*[]DashboardReservation2, error)
}
