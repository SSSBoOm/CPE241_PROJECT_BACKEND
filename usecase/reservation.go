package usecase

import (
	"sync"
	"time"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
)

type reservationUsecase struct {
	reservationRepository domain.ReservationRepository
	roomTypeUsecase       domain.RoomTypeUsecase
	roomUsecase           domain.RoomUsecase
}

func NewReservationUsecase(reservationRepository domain.ReservationRepository, roomTypeUsecase domain.RoomTypeUsecase, roomUsecase domain.RoomUsecase) domain.ReservationUsecase {
	return &reservationUsecase{
		reservationRepository: reservationRepository,
		roomTypeUsecase:       roomTypeUsecase,
		roomUsecase:           roomUsecase,
	}
}

func (u *reservationUsecase) GetAll() (*[]domain.RESERVATION, error) {
	return u.reservationRepository.GetAll()
}

func (u *reservationUsecase) GetByDate(startDate string, endDate string) (*[]domain.RESERVATION, error) {
	return u.reservationRepository.GetByDate(startDate, endDate)
}

func (u *reservationUsecase) GetByUserID(userID string) (*[]domain.RESERVATION, error) {
	return u.reservationRepository.GetByUserID(userID)
}

func (u *reservationUsecase) GetByRoomID(roomID int) (*[]domain.RESERVATION, error) {
	return u.reservationRepository.GetByRoomID(roomID)
}

func (u *reservationUsecase) GetByID(id int) (*domain.RESERVATION, error) {
	return u.reservationRepository.GetByID(id)
}

func (u *reservationUsecase) Create(reservation *domain.RESERVATION) (id *int, err error) {
	return u.reservationRepository.Create(reservation)
}

func (u *reservationUsecase) Update(reservation *domain.RESERVATION) error {
	return u.reservationRepository.Update(reservation)
}

func (u *reservationUsecase) GetRoomAvailableGroupByRoomType(start_date time.Time, end_date time.Time) ([]domain.RoomType, error) {
	roomType, err := u.roomTypeUsecase.GetAll()
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	wg.Add(len(*roomType))
	// for _, item := range *roomType {
	// 	go func(i int, item domain.Payment) {
	// 		defer wg.Done()
	// 	}
	// }
	// wg.Wait()

	return nil, nil
}
