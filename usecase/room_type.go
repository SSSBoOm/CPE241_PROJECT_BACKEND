package usecase

import "github.com/SSSBoOm/CPE241_Project_Backend/domain"

type roomTypeUsecase struct {
	roomTypeRepository domain.RoomTypeRepository
}

func NewRoomTypeUsecase(roomTypeRepository domain.RoomTypeRepository) domain.RoomTypeUsecase {
	return &roomTypeUsecase{
		roomTypeRepository: roomTypeRepository,
	}
}

func (u *roomTypeUsecase) GetAll() (*[]domain.RoomType, error) {
	return u.roomTypeRepository.GetAll()
}

func (u *roomTypeUsecase) GetByID(id int) (*domain.RoomType, error) {
	return u.roomTypeRepository.GetByID(id)
}

func (u *roomTypeUsecase) Create(roomType *domain.RoomType) error {
	return u.roomTypeRepository.Create(roomType)
}

func (u *roomTypeUsecase) Update(roomType *domain.RoomType) error {
	return u.roomTypeRepository.Update(roomType)
}
