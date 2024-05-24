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
	serviceUsecase        domain.ServiceUsecase
	paymentUsecase        domain.PaymentUsecase
	userUsecase           domain.UserUsecase
}

func NewReservationUsecase(reservationRepository domain.ReservationRepository, roomTypeUsecase domain.RoomTypeUsecase, roomUsecase domain.RoomUsecase, serviceUsecase domain.ServiceUsecase, paymentUsecase domain.PaymentUsecase, userUsecase domain.UserUsecase) domain.ReservationUsecase {
	return &reservationUsecase{
		reservationRepository: reservationRepository,
		roomTypeUsecase:       roomTypeUsecase,
		roomUsecase:           roomUsecase,
		serviceUsecase:        serviceUsecase,
		paymentUsecase:        paymentUsecase,
		userUsecase:           userUsecase,
	}
}

func (u *reservationUsecase) GetAll() (*[]domain.RESERVATION, error) {
	reservation, err := u.reservationRepository.GetAll()
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	wg.Add(len(*reservation))
	for i, item := range *reservation {
		go func(i int, item domain.RESERVATION) {
			defer wg.Done()
			if item.TYPE == domain.RESERVATION_TYPE_ROOM {
				room, err := u.roomUsecase.GetByID(*item.ROOM_ID)
				(*reservation)[i].ROOM = room
				if err != nil {
					return
				}
			} else if item.TYPE == domain.RESERVATION_TYPE_SERVICE {
				service, err := u.serviceUsecase.GetById(*item.SERVICE_ID)
				(*reservation)[i].SERVICE = service
				if err != nil {
					return
				}
			}

			payment, err := u.paymentUsecase.GetByID(item.PAYMENT_INFO_ID)
			(*reservation)[i].PAYMENT_INFO = payment
			if err != nil {
				return
			}

			if item.STAFF_ID != nil {
				staff, err := u.userUsecase.FindById(*item.STAFF_ID)
				(*reservation)[i].STAFF = staff
				if err != nil {
					return
				}
			}

			if item.USER_ID != "" {
				user, err := u.userUsecase.FindById(item.USER_ID)
				(*reservation)[i].USER = user
				if err != nil {
					return
				}
			}
		}(i, item)
	}

	wg.Wait()
	return reservation, nil
}

func (u *reservationUsecase) GetByDate(startDate string, endDate string) (*[]domain.RESERVATION, error) {
	return u.reservationRepository.GetByDate(startDate, endDate)
}

func (u *reservationUsecase) GetByUserID(userID string) (*[]domain.RESERVATION, error) {
	reservation, err := u.reservationRepository.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	wg.Add(len(*reservation))
	for i, item := range *reservation {
		go func(i int, item domain.RESERVATION) {
			defer wg.Done()
			if item.TYPE == domain.RESERVATION_TYPE_ROOM {
				room, err := u.roomUsecase.GetByID(*item.ROOM_ID)
				(*reservation)[i].ROOM = room
				if err != nil {
					return
				}
			} else if item.TYPE == domain.RESERVATION_TYPE_SERVICE {
				service, err := u.serviceUsecase.GetById(*item.SERVICE_ID)
				(*reservation)[i].SERVICE = service
				if err != nil {
					return
				}
			}

			payment, err := u.paymentUsecase.GetByID(item.PAYMENT_INFO_ID)
			(*reservation)[i].PAYMENT_INFO = payment
			if err != nil {
				return
			}

			if item.STAFF_ID != nil {
				staff, err := u.userUsecase.FindById(*item.STAFF_ID)
				(*reservation)[i].STAFF = staff
				if err != nil {
					return
				}
			}
		}(i, item)
	}
	wg.Wait()

	return reservation, nil
}

func (u *reservationUsecase) GetByRoomID(roomID int) (*[]domain.RESERVATION, error) {
	return u.reservationRepository.GetByRoomID(roomID)
}

func (u *reservationUsecase) GetByID(id int) (*domain.RESERVATION, error) {
	reservation, err := u.reservationRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	if reservation.TYPE == domain.RESERVATION_TYPE_ROOM {
		room, err := u.roomUsecase.GetByID(*reservation.ROOM_ID)
		reservation.ROOM = room
		if err != nil {
			return nil, err
		}
	} else if reservation.TYPE == domain.RESERVATION_TYPE_SERVICE {
		service, err := u.serviceUsecase.GetById(*reservation.SERVICE_ID)
		reservation.SERVICE = service
		if err != nil {
			return nil, err
		}
	}
	payment, err := u.paymentUsecase.GetByID(reservation.PAYMENT_INFO_ID)
	reservation.PAYMENT_INFO = payment
	if err != nil {
		return nil, err
	}
	if reservation.STAFF_ID != nil {
		staff, err := u.userUsecase.FindById(*reservation.STAFF_ID)
		reservation.STAFF = staff
		if err != nil {
			return nil, err
		}
	}
	if reservation.USER_ID != "" {
		user, err := u.userUsecase.FindById(reservation.USER_ID)
		reservation.USER = user
		if err != nil {
			return nil, err
		}
	}

	return reservation, nil
}

func (u *reservationUsecase) Create(reservation *domain.RESERVATION) (id *int, err error) {
	return u.reservationRepository.Create(reservation)
}

func (u *reservationUsecase) Update(reservation *domain.RESERVATION) error {
	return u.reservationRepository.Update(reservation)
}

func (u *reservationUsecase) UpdateStaff(id int, staffID string) error {
	return u.reservationRepository.UpdateStaff(id, staffID)
}

func (u *reservationUsecase) UpdateStatus(id int, status domain.RESERVATION_STATUS) error {
	return u.reservationRepository.UpdateStatus(id, status)
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

func (u *reservationUsecase) UpdatePayment(id int, paymentInfoID int) error {
	return u.reservationRepository.UpdatePayment(id, paymentInfoID)
}
