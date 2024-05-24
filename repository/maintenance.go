package repository

import (
	"time"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/jmoiron/sqlx"
)

type maintenanceRepository struct {
	db *sqlx.DB
}

func NewMaintenanceRepository(db *sqlx.DB) domain.MaintenanceRepository {
	return &maintenanceRepository{db: db}
}

func (repo *maintenanceRepository) GetByID(id int) (*domain.MAINTENANCE, error) {
	maintenance := domain.MAINTENANCE{}
	err := repo.db.Get(&maintenance, "SELECT * FROM maintenance WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &maintenance, nil
}

func (repo *maintenanceRepository) GetAll() (*[]domain.MAINTENANCE, error) {
	maintenances := make([]domain.MAINTENANCE, 0)
	err := repo.db.Select(&maintenances, "SELECT * FROM maintenance")
	if err != nil {
		return nil, err
	}
	return &maintenances, nil
}

func (repo *maintenanceRepository) Create(maintenance *domain.MAINTENANCE) (*int, error) {
	t := repo.db.MustBegin()
	data, err := t.NamedExec("INSERT INTO maintenance (title, room_id, staff_id) VALUES (:title, :room_id, :staff_id)", maintenance)
	if err != nil {
		t.Rollback()
		return nil, err
	}
	t.Commit()
	LastInsertId, _ := data.LastInsertId()
	Id := int(LastInsertId)
	return &Id, nil
}

func (repo *maintenanceRepository) Update(maintenance *domain.MAINTENANCE) error {
	maintenance.UPDATED_AT = time.Now()
	t := repo.db.MustBegin()
	_, err := t.NamedExec("UPDATE maintenance SET title = :title, room_id = :room_id, staff_id = :staff_id, updated_at := updated_at WHERE id = :id", maintenance)
	if err != nil {
		t.Rollback()
		return err
	}
	t.Commit()
	return nil
}
