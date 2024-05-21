package usecase

import "github.com/SSSBoOm/CPE241_Project_Backend/domain"

type reservationTaskUseCase struct {
	reservationTaskRepository domain.ReservationTaskRepository
}

func NewReservationTaskUseCase(reservationTaskRepository domain.ReservationTaskRepository) domain.ReservationTaskUsecase {
	return &reservationTaskUseCase{
		reservationTaskRepository: reservationTaskRepository,
	}
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
