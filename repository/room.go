package repository

import (
	"time"

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

func (r *roomRepository) GetAll() (*[]domain.ROOM, error) {
	room := make([]domain.ROOM, 0)
	err := r.db.Select(&room, "SELECT * FROM room")
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *roomRepository) GetByRoomType(roomTypeID int) (*[]domain.ROOM, error) {
	room := make([]domain.ROOM, 0)
	err := r.db.Select(&room, "SELECT * FROM room WHERE room_type_id = ?", roomTypeID)
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *roomRepository) GetByID(id int) (*domain.ROOM, error) {
	room := domain.ROOM{}
	err := r.db.Get(&room, "SELECT * FROM room WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *roomRepository) Create(room *domain.ROOM) error {
	t := r.db.MustBegin()
	_, err := t.NamedExec("INSERT INTO room (room_number, is_active, room_type_id) VALUES (:room_number, :is_active, :room_type_id)", room)
	if err != nil {
		t.Rollback()
		return err
	}
	t.Commit()
	return nil
}

func (r *roomRepository) Update(room *domain.ROOM) error {
	room.UPDATED_AT = time.Now()
	t := r.db.MustBegin()
	_, err := t.NamedExec("UPDATE room SET room_number = :room_number, is_active = :is_active, room_type_id = :room_type_id, update_at = :update_at WHERE id = :id", room)
	if err != nil {
		t.Rollback()
		return err
	}
	t.Commit()
	return nil
}

func (r *roomRepository) UpdateIsActive(id int, isActive bool) error {
	t := r.db.MustBegin()
	_, err := t.Exec("UPDATE room SET is_active = ?, updated_at = ? WHERE id = ?", isActive, time.Now(), id)
	if err != nil {
		t.Rollback()
		return err
	}
	t.Commit()
	return nil
}
