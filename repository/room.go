package repository

import (
	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/jmoiron/sqlx"
)

type roomRepository struct {
	db *sqlx.DB
}

func NewRoomRepository(db *sqlx.DB) domain.RoomRepository {
	return &roomRepository{
		db: db,
	}
}

func (r *roomRepository) GetAll() (*[]domain.Room, error) {
	room := make([]domain.Room, 0)
	err := r.db.Select(&room, "SELECT * FROM room_type")
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *roomRepository) GetByID(id int) (*domain.Room, error) {
	room := domain.Room{}
	err := r.db.Get(&room, "SELECT * FROM room WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *roomRepository) Create(room *domain.Room) error {
	t := r.db.MustBegin()
	_, err := t.NamedExec("INSERT INTO room (room_number, is_active, room_type_id) VALUES (:room_number, :is_active, :room_type_id)", room)
	if err != nil {
		t.Rollback()
		return err
	}
	t.Commit()
	return nil
}

func (r *roomRepository) Update(room *domain.Room) error {
	t := r.db.MustBegin()
	_, err := t.NamedExec("UPDATE room SET room_number = :room_number, is_active = :is_active, room_type_id = :room_type_id, update_at = :update_at WHERE id = :id", room)
	if err != nil {
		t.Rollback()
		return err
	}
	t.Commit()
	return nil
}
