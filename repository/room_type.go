package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/jmoiron/sqlx"
)

type roomTypeRepository struct {
	db *sqlx.DB
}

func NewRoomTypeRepository(db *sqlx.DB) domain.RoomTypeRepository {
	return &roomTypeRepository{
		db: db,
	}
}

func (r *roomTypeRepository) GetAll() (*[]domain.RoomType, error) {
	roomType := make([]domain.RoomType, 0)
	err := r.db.Select(&roomType, "SELECT * FROM room_type")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &roomType, nil
}

func (r *roomTypeRepository) GetByID(id int) (*domain.RoomType, error) {
	roomType := domain.RoomType{}
	err := r.db.Get(&roomType, "SELECT * FROM room_type WHERE id = ?", id)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &roomType, nil
}

func (r *roomTypeRepository) Create(roomType *domain.RoomType) error {
	t := r.db.MustBegin()
	_, err := t.NamedExec("INSERT INTO room_type (name, detail, is_active) VALUES (:name, :detail, :is_active)", roomType)
	if err != nil {
		t.Rollback()
		return err
	}
	t.Commit()
	return nil
}

func (r *roomTypeRepository) Update(roomType *domain.RoomType) error {
	roomType.UPDATED_AT = time.Now()
	t := r.db.MustBegin()
	_, err := t.NamedExec("UPDATE room_type SET name = :name, detail = :detail, is_active = :is_active, update_at = :update_at WHERE id = :id", roomType)
	if err != nil {
		t.Rollback()
		return err
	}
	t.Commit()
	return nil
}