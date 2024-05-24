package usecase

import (
	"sync"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
)

type maintenanceUsecase struct {
	maintenanceRepository domain.MaintenanceRepository
	maintenanceLogUsecase domain.MaintenanceLogUsecase
	roomUsecase           domain.RoomUsecase
}

func NewMaintenanceUsecase(maintenanceRepository domain.MaintenanceRepository, maintenanceLogUsecase domain.MaintenanceLogUsecase, roomUsecase domain.RoomUsecase) domain.MaintenanceUsecase {
	return &maintenanceUsecase{
		maintenanceRepository: maintenanceRepository,
		maintenanceLogUsecase: maintenanceLogUsecase,
		roomUsecase:           roomUsecase,
	}
}

func (u *maintenanceUsecase) GetByID(id int) (*domain.MAINTENANCE, error) {
	maintenance, err := u.maintenanceRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	room, err := u.roomUsecase.GetByID(maintenance.ROOM_ID)
	if err != nil {
		return nil, err
	}
	maintenance.ROOM = room

	maintenanceLogs, err := u.maintenanceLogUsecase.GetByMaintenanceID(id)
	if err != nil {
		return nil, err
	}
	maintenance.MAINTENANCE_LOG = maintenanceLogs
	return maintenance, nil
}

func (u *maintenanceUsecase) GetAll() (*[]domain.MAINTENANCE, error) {
	maintenance, err := u.maintenanceRepository.GetAll()
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	wg.Add(len(*maintenance))
	for i, item := range *maintenance {
		go func(i int, item domain.MAINTENANCE) {
			defer wg.Done()
			room, err := u.roomUsecase.GetByID(item.ROOM_ID)
			if err != nil {
				return
			}
			(*maintenance)[i].ROOM = room

			maintenanceLogs, err := u.maintenanceLogUsecase.GetByMaintenanceID(item.ID)
			if err != nil {
				return
			}
			(*maintenance)[i].MAINTENANCE_LOG = maintenanceLogs
		}(i, item)
	}
	wg.Wait()

	return maintenance, nil
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

	var wg sync.WaitGroup
	wg.Add(len(*maintenance.MAINTENANCE_LOG))
	for i, item := range *maintenance.MAINTENANCE_LOG {
		go func(i int, item domain.MAINTENANCE_LOG) {
			defer wg.Done()
			item.MAINTENANCE_ID = *id
			item.STAFF_ID = maintenance.STAFF_ID
			err = u.maintenanceLogUsecase.Create(&item)
			if err != nil {
				return
			}
		}(i, item)
	}
	wg.Wait()

	return nil
}

func (u *maintenanceUsecase) Update(maintenance *domain.MAINTENANCE) error {
	return u.maintenanceRepository.Update(maintenance)
}
