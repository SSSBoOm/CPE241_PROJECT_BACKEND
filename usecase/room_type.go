package usecase

import (
	"sync"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
)

type roomTypeUsecase struct {
	roomTypeRepository domain.RoomTypeRepository
	roomRepository     domain.RoomRepository
}

func NewRoomTypeUsecase(roomTypeRepository domain.RoomTypeRepository, roomRepository domain.RoomRepository) domain.RoomTypeUsecase {
	return &roomTypeUsecase{
		roomTypeRepository: roomTypeRepository,
		roomRepository:     roomRepository,
	}
}

func (u *roomTypeUsecase) GetAll() (*[]domain.RoomType, error) {
	roomType, err := u.roomTypeRepository.GetAll()
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	wg.Add(len(*roomType))
	for i, item := range *roomType {
		go func(i int, item domain.RoomType) {
			defer wg.Done()
			rooms, err := u.roomRepository.GetByRoomType(item.ID)
			if err != nil {
				return
			}
			(*roomType)[i].ROOM = rooms
		}(i, item)
	}
	wg.Wait()

	return roomType, nil
}

func (u *roomTypeUsecase) GetByID(id int) (*domain.RoomType, error) {
	return u.roomTypeRepository.GetByID(id)
}

func (u *roomTypeUsecase) Create(roomType *domain.RoomType) (*int, error) {
	roomTypeId, err := u.roomTypeRepository.Create(roomType)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	wg.Add(len(*roomType.ROOM))
	for i, item := range *roomType.ROOM {
		go func(i int, item domain.ROOM) {
			defer wg.Done()
			item.ROOM_TYPE_ID = *roomTypeId
			err = u.roomRepository.Create(&item)
			if err != nil {
				return
			}
		}(i, item)
	}
	wg.Wait()

	return roomTypeId, nil
}

func (u *roomTypeUsecase) Update(roomType *domain.RoomType) error {
	return u.roomTypeRepository.Update(roomType)
}

func (u *roomTypeUsecase) UpdateIsActive(id int, isActive bool) error {
	return u.roomTypeRepository.UpdateIsActive(id, isActive)
}
