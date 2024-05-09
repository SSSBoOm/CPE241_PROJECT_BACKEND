package usecase

import "github.com/SSSBoOm/CPE241_Project_Backend/domain"

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
	return u.roomRepository.GetAll()
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
