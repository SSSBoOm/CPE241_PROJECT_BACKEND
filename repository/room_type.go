package repository

import (
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
		return nil, err
	}
	return &roomType, nil
}

func (r *roomTypeRepository) GetByID(id int) (*domain.RoomType, error) {
	roomType := domain.RoomType{}
	err := r.db.Get(&roomType, "SELECT * FROM room_type WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &roomType, nil
}

func (r *roomTypeRepository) Create(roomType *domain.RoomType) error {
	_, err := r.db.NamedExec("INSERT INTO room_type (name, is_active) VALUES (:name, :is_active)", roomType)
	if err != nil {
		return err
	}
	return nil
}

func (r *roomTypeRepository) Update(roomType *domain.RoomType) error {
	_, err := r.db.NamedExec("UPDATE room_type SET name = :name, is_active = :is_active WHERE id = :id", roomType)
	if err != nil {
		return err
	}
	return nil
}
