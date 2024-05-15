package usecase

import (
	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
)

type maintenanceUsecase struct {
	maintenanceRepo       domain.MaintenanceRepository
	maintenanceLogUsecase domain.MaintenanceLogUsecase
}

func NewMaintenanceUsecase(maintenanceRepo domain.MaintenanceRepository, maintenanceLogUsecase domain.MaintenanceLogUsecase) domain.MaintenanceUsecase {
	return &maintenanceUsecase{maintenanceRepo: maintenanceRepo,
		maintenanceLogUsecase: maintenanceLogUsecase,
	}
}

func (u *maintenanceUsecase) Get(id int) (*domain.MAINTENANCE, error) {
	maintenance, err := u.maintenanceRepo.Get(id)
	if err != nil {
		return nil, err
	}
	maintenanceLogs, err := u.maintenanceLogUsecase.GetByMaintenanceID(id)
	if err != nil {
		return nil, err
	}
	maintenance.MAINTENANCE_LOG = maintenanceLogs
	return maintenance, nil
}

func (u *maintenanceUsecase) GetAll() (*[]domain.MAINTENANCE, error) {
	return u.maintenanceRepo.GetAll()
}

func (u *maintenanceUsecase) Create(maintenance *domain.MAINTENANCE) error {
	return u.maintenanceRepo.Create(maintenance)
}

func (u *maintenanceUsecase) Update(maintenance *domain.MAINTENANCE) error {
	return u.maintenanceRepo.Update(maintenance)
}
