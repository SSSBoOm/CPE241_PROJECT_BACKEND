package domain

import "time"

type DashboardReservation struct {
	ID    int    `json:"id" db:"id"`
	NAME  string `json:"name" db:"name"`
	TOTAL int    `json:"total" db:"total"`
}

type DashboardService struct {
}

type DashboardUsecase interface {
	GetDashboardReservation(startDate time.Time, endDate time.Time) (*[]DashboardReservation, error)
}

type DashboardRepository interface {
	GetDashboardReservation(startDate time.Time, endDate time.Time) (*[]DashboardReservation, error)
}
