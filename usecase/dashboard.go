package usecase

import (
	"time"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
)

type dashboardUsecase struct {
	dashboardRepository domain.DashboardRepository
}

func NewDashboardUsecase(dashboardRepository domain.DashboardRepository) domain.DashboardUsecase {
	return &dashboardUsecase{
		dashboardRepository: dashboardRepository,
	}
}

func (u *dashboardUsecase) GetDashboardReservation(startDate time.Time, endDate time.Time) (*[]domain.DashboardReservation, error) {
	return u.dashboardRepository.GetDashboardReservation(startDate, endDate)
}
