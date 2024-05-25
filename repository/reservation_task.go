package repository

import (
	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/jmoiron/sqlx"
)

type reservationTaskRepository struct {
	db *sqlx.DB
}

func NewReservationTaskRepository(db *sqlx.DB) domain.ReservationTaskRepository {
	return &reservationTaskRepository{db: db}
}

func (r *reservationTaskRepository) GetAll() (*[]domain.RESERVATION_TASK, error) {
	tasks := make([]domain.RESERVATION_TASK, 0)
	err := r.db.Select(&tasks, "SELECT * FROM reservation_task")
	if err != nil {
		return nil, err
	}
	return &tasks, nil
}

func (r *reservationTaskRepository) Create(task *domain.RESERVATION_TASK) error {
	t := r.db.MustBegin()
	_, err := t.NamedExec("INSERT INTO reservation_task (reservation_id, status, date) VALUES (:reservation_id, :status, :date)", task)
	if err != nil {
		t.Rollback()
		return err
	}
	t.Commit()
	return nil
}

func (r *reservationTaskRepository) GetByReservationID(reservationID int) (*domain.RESERVATION_TASK, error) {
	var task domain.RESERVATION_TASK
	err := r.db.Get(&task, "SELECT * FROM reservation_task WHERE reservation_id = $1", reservationID)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *reservationTaskRepository) Update(task *domain.RESERVATION_TASK) error {
	t := r.db.MustBegin()
	_, err := t.NamedExec("UPDATE reservation_task SET staff_id = :staff_id, status = :status, date = :date WHERE id = :id", task)
	if err != nil {
		t.Rollback()
		return err
	}
	t.Commit()
	return nil
}

func (r *reservationTaskRepository) UpdateStaff(id int, staffID string) error {
	t := r.db.MustBegin()
	_, err := t.Exec("UPDATE reservation_task SET staff_id = ? WHERE id = ?", staffID, id)
	if err != nil {
		t.Rollback()
		return err
	}
	t.Commit()
	return nil
}

func (r *reservationTaskRepository) UpdateStatus(id int, status bool) error {
	t := r.db.MustBegin()
	_, err := t.Exec("UPDATE reservation_task SET status = ? WHERE id = ?", status, id)
	if err != nil {
		t.Rollback()
		return err
	}
	t.Commit()
	return nil
}
