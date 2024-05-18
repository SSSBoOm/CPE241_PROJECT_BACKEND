package usecase

import (
	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
)

type maintenanceUsecase struct {
	maintenanceRepository domain.MaintenanceRepository
	maintenanceLogUsecase domain.MaintenanceLogUsecase
}

func NewMaintenanceUsecase(maintenanceRepository domain.MaintenanceRepository, maintenanceLogUsecase domain.MaintenanceLogUsecase) domain.MaintenanceUsecase {
	return &maintenanceUsecase{maintenanceRepository: maintenanceRepository,
		maintenanceLogUsecase: maintenanceLogUsecase,
	}
}

func (u *maintenanceUsecase) GetByID(id int) (*domain.MAINTENANCE, error) {
	maintenance, err := u.maintenanceRepository.GetByID(id)
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
	return u.maintenanceRepository.GetAll()
}

func (u *maintenanceUsecase) Create(maintenance *domain.MAINTENANCE) error {
	_, err := u.maintenanceRepository.Create(maintenance)
	return err
}

func (u *maintenanceUsecase) CreateWithMaintenance_Log(maintenance *domain.MAINTENANCE) error {
	id, err := u.maintenanceRepository.Create(maintenance)
	if err != nil {
		return err
	}
	for _, maintenanceLog := range *maintenance.MAINTENANCE_LOG {
		maintenanceLog.MAINTENANCE_ID = *id
		maintenanceLog.STAFF_ID = maintenance.STAFF_ID
		err = u.maintenanceLogUsecase.Create(&maintenanceLog)
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *maintenanceUsecase) Update(maintenance *domain.MAINTENANCE) error {
	return u.maintenanceRepository.Update(maintenance)
}
