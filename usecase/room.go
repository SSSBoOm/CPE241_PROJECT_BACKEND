package usecase

import (
	"sync"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
)

type RoomUsecase struct {
	roomRepository  domain.RoomRepository
	roomTypeUsecase domain.RoomTypeUsecase
}

func NewRoomUsecase(roomRepository domain.RoomRepository, roomTypeUsecase domain.RoomTypeUsecase) domain.RoomUsecase {
	return &RoomUsecase{
		roomRepository:  roomRepository,
		roomTypeUsecase: roomTypeUsecase,
	}
}

func (u *RoomUsecase) GetAll() (*[]domain.ROOM, error) {
	room, err := u.roomRepository.GetAll()
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	wg.Add(len(*room))
	for i, item := range *room {
		go func(i int, item domain.ROOM) {
			defer wg.Done()
			roomType, err := u.roomTypeUsecase.GetByID(item.ROOM_TYPE_ID)
			if err != nil || roomType == nil {
				return
			}
			(*room)[i].ROOM_TYPE = roomType
		}(i, item)
	}
	wg.Wait()

	return room, nil
}

func (u *RoomUsecase) GetByID(id int) (*domain.ROOM, error) {
	room, err := u.roomRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	roomType, err := u.roomTypeUsecase.GetByID(room.ROOM_TYPE_ID)
	if err != nil || roomType == nil {
		return room, nil
	}
	room.ROOM_TYPE = roomType
	return room, nil
}

func (u *RoomUsecase) GetByRoomType(roomTypeID int) (*[]domain.ROOM, error) {
	return u.roomRepository.GetByRoomType(roomTypeID)
}

func (u *RoomUsecase) Create(room *domain.ROOM) error {
	return u.roomRepository.Create(room)
}

func (u *RoomUsecase) Update(room *domain.ROOM) error {
	return u.roomRepository.Update(room)
}

func (u *RoomUsecase) UpdateIsActive(id int, isActive bool) error {
	return u.roomRepository.UpdateIsActive(id, isActive)
}
