package usecase

import (
	"sync"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
)

type maintenanceLogUsecase struct {
	maintenanceLogRepo domain.MaintenanceLogRepository
	userUsecase        domain.UserUsecase
}

func NewMaintenanceLogUsecase(maintenanceLogRepo domain.MaintenanceLogRepository, userUsecase domain.UserUsecase) domain.MaintenanceLogUsecase {
	return &maintenanceLogUsecase{
		maintenanceLogRepo: maintenanceLogRepo,
		userUsecase:        userUsecase,
	}
}

func (u *maintenanceLogUsecase) GetByID(id int) (*domain.MAINTENANCE_LOG, error) {
	maintenanceLog, err := u.maintenanceLogRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	user, err := u.userUsecase.FindById(maintenanceLog.STAFF_ID)
	if err != nil {
		return nil, err
	}

	maintenanceLog.STAFF = user
	return maintenanceLog, nil
}

func (u *maintenanceLogUsecase) GetByMaintenanceID(maintenance_id int) (*[]domain.MAINTENANCE_LOG, error) {
	maintenanceLog, err := u.maintenanceLogRepo.GetByMaintenanceID(maintenance_id)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	wg.Add(len(*maintenanceLog))
	for i, item := range *maintenanceLog {
		go func(i int, item domain.MAINTENANCE_LOG) {
			defer wg.Done()
			user, err := u.userUsecase.FindById(item.STAFF_ID)
			if err != nil {
				return
			}
			item.STAFF = user
			(*maintenanceLog)[i] = item
		}(i, item)
	}
	wg.Wait()

	return maintenanceLog, nil
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
