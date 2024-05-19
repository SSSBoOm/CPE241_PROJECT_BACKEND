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

func (u *RoomUsecase) GetAll() (*[]domain.Room, error) {
	room, err := u.roomRepository.GetAll()
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	wg.Add(len(*room))
	for i, item := range *room {
		go func(i int, item domain.Room) {
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

func (u *RoomUsecase) GetByID(id int) (*domain.Room, error) {
	return u.roomRepository.GetByID(id)
}

func (u *RoomUsecase) Create(room *domain.Room) error {
	return u.roomRepository.Create(room)
}

func (u *RoomUsecase) Update(room *domain.Room) error {
	return u.roomRepository.Update(room)
}
