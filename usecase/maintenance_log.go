package usecase

import "github.com/SSSBoOm/CPE241_Project_Backend/domain"

type maintenanceLogUsecase struct {
	maintenanceLogRepo domain.MaintenanceLogRepository
}

func NewMaintenanceLogUsecase(maintenanceLogRepo domain.MaintenanceLogRepository) domain.MaintenanceLogUsecase {
	return &maintenanceLogUsecase{maintenanceLogRepo: maintenanceLogRepo}
}

func (u *maintenanceLogUsecase) GetByID(id int) (*domain.MAINTENANCE_LOG, error) {
	return u.maintenanceLogRepo.GetByID(id)
}

func (u *maintenanceLogUsecase) GetByMaintenanceID(maintenance_id int) (*[]domain.MAINTENANCE_LOG, error) {
	return u.maintenanceLogRepo.GetByMaintenanceID(maintenance_id)
}

func (u *maintenanceLogUsecase) GetAll() (*[]domain.MAINTENANCE_LOG, error) {
	return u.maintenanceLogRepo.GetAll()
}

func (u *maintenanceLogUsecase) Create(maintenanceLog *domain.MAINTENANCE_LOG) error {
	return u.maintenanceLogRepo.Create(maintenanceLog)
}

func (u *maintenanceLogUsecase) Update(maintenanceLog *domain.MAINTENANCE_LOG) error {
	return u.maintenanceLogRepo.Update(maintenanceLog)
}
