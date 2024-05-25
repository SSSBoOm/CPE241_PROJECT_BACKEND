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

func (u *dashboardUsecase) GetDashboardRoomTypeReservation(startDate time.Time, endDate time.Time) (*[]domain.DashboardReservation, error) {
	return u.dashboardRepository.GetDashboardRoomTypeReservation(startDate, endDate)
}

func (u *dashboardUsecase) GetRoomTypeReservationCountByBooking(startDate time.Time, endDate time.Time) (*[]domain.DashboardReservation, error) {
	return u.dashboardRepository.GetRoomTypeReservationCountByBooking(startDate, endDate)
}

func (u *dashboardUsecase) GetDashboardServiceTypeReservation(startDate time.Time, endDate time.Time) (*[]domain.DashboardReservation, error) {
	return u.dashboardRepository.GetDashboardServiceTypeReservation(startDate, endDate)
}

func (u *dashboardUsecase) GetServiceTypeReservationCountByBooking(startDate time.Time, endDate time.Time) (*[]domain.DashboardReservation, error) {
	return u.dashboardRepository.GetServiceTypeReservationCountByBooking(startDate, endDate)
}

func (u *dashboardUsecase) GetDashboardReservationByPaymentType(startDate time.Time, endDate time.Time) (*[]domain.DashboardReservation2, error) {
	return u.dashboardRepository.GetDashboardReservationByPaymentType(startDate, endDate)
}

func (u *dashboardUsecase) GetTotalMaintenanceByRoomType() (*[]domain.DashboardReservation, error) {
	return u.dashboardRepository.GetTotalMaintenanceByRoomType()
}
