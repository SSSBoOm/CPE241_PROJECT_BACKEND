package repository

import (
	"time"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/jmoiron/sqlx"
)

type reservationRepository struct {
	db *sqlx.DB
}

func NewReservationRepository(db *sqlx.DB) domain.ReservationRepository {
	return &reservationRepository{
		db: db,
	}
}

func (repo *reservationRepository) GetAll() (*[]domain.RESERVATION, error) {
	reservations := make([]domain.RESERVATION, 0)
	err := repo.db.Select(&reservations, "SELECT * FROM reservation")
	if err != nil {
		return nil, err
	}
	return &reservations, nil
}

func (repo *reservationRepository) GetByDate(startDate string, endDate string) (*[]domain.RESERVATION, error) {
	reservations := make([]domain.RESERVATION, 0)
	err := repo.db.Select(&reservations, "SELECT * FROM reservation WHERE start_date >= ? AND end_date <= ?", startDate, endDate)
	if err != nil {
		return nil, err
	}
	return &reservations, nil
}

func (repo *reservationRepository) GetByUserID(userID string) (*[]domain.RESERVATION, error) {
	reservations := make([]domain.RESERVATION, 0)
	err := repo.db.Select(&reservations, "SELECT * FROM reservation WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	return &reservations, nil
}

func (repo *reservationRepository) GetByRoomID(roomID int) (*[]domain.RESERVATION, error) {
	reservations := make([]domain.RESERVATION, 0)
	err := repo.db.Select(&reservations, "SELECT * FROM reservation WHERE room_id = ?", roomID)
	if err != nil {
		return nil, err
	}
	return &reservations, nil
}

func (repo *reservationRepository) GetByID(id int) (*domain.RESERVATION, error) {
	reservation := domain.RESERVATION{}
	err := repo.db.Get(&reservation, "SELECT * FROM reservation WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &reservation, nil
}

func (repo *reservationRepository) Create(reservation *domain.RESERVATION) (id *int, err error) {
	t := repo.db.MustBegin()
	reservation.STATUS = domain.RESERVATION_STATUS_WAITING_APPROVE_PAYMENT
	result, err := t.NamedExec("INSERT INTO reservation (room_id, user_id, start_date, end_date, price, status, payment_date, payment_info_id) VALUES (:room_id, :user_id, :start_date, :end_date, :price, :status, :payment_date, :payment_info_id)", reservation)
	if err != nil {
		t.Rollback()
		return nil, err
	}
	t.Commit()
	LastInsertId, _ := result.LastInsertId()
	Id := int(LastInsertId)
	return &Id, nil
}

func (repo *reservationRepository) Update(reservation *domain.RESERVATION) error {
	reservation.UPDATED_AT = time.Now()
	t := repo.db.MustBegin()
	_, err := t.NamedExec("UPDATE reservation SET room_id = :room_id, user_id = :user_id, start_date = :start_date, end_date = :end_date, status = :status, updated_at = :updated_at WHERE id = :id", reservation)
	if err != nil {
		t.Rollback()
		return err
	}
	t.Commit()
	return nil
}
