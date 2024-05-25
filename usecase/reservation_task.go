package usecase

import (
	"sync"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
)

type reservationTaskUseCase struct {
	reservationTaskRepository domain.ReservationTaskRepository
	reservationUsecase        domain.ReservationUsecase
	userUsecase               domain.UserUsecase
}

func NewReservationTaskUseCase(reservationTaskRepository domain.ReservationTaskRepository, reservationUsecase domain.ReservationUsecase, userUsecase domain.UserUsecase) domain.ReservationTaskUsecase {
	return &reservationTaskUseCase{
		reservationTaskRepository: reservationTaskRepository,
		reservationUsecase:        reservationUsecase,
		userUsecase:               userUsecase,
	}
}

func (u *reservationTaskUseCase) GetAll() (*[]domain.RESERVATION_TASK, error) {
	task, err := u.reservationTaskRepository.GetAll()
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	wg.Add(len(*task))
	for i := range *task {
		go func(i int) {
			defer wg.Done()
			if (*task)[i].STAFF_ID != nil {
				staff, err := u.userUsecase.FindById(*(*task)[i].STAFF_ID)
				if err != nil {
					return
				}
				(*task)[i].STAFF = staff
			}

			reservation, err := u.reservationUsecase.GetByID((*task)[i].RESERVATION_ID)
			if err != nil {
				return
			}
			(*task)[i].RESERVATION = reservation
		}(i)
	}
	wg.Wait()

	return task, nil
}

func (u *reservationTaskUseCase) Create(task *domain.RESERVATION_TASK) error {
	return u.reservationTaskRepository.Create(task)
}

func (u *reservationTaskUseCase) GetByReservationID(reservationID int) (*domain.RESERVATION_TASK, error) {
	return u.reservationTaskRepository.GetByReservationID(reservationID)
}

func (u *reservationTaskUseCase) Update(task *domain.RESERVATION_TASK) error {
	return u.reservationTaskRepository.Update(task)
}

func (u *reservationTaskUseCase) UpdateStaff(id int, staffID string) error {
	return u.reservationTaskRepository.UpdateStaff(id, staffID)
}

func (u *reservationTaskUseCase) UpdateStatus(id int, status bool) error {
	return u.reservationTaskRepository.UpdateStatus(id, status)
}
