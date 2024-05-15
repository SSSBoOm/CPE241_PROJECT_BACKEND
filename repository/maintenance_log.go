package repository

import (
	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/jmoiron/sqlx"
)

type maintenanceLogRepository struct {
	db *sqlx.DB
}

func NewMaintenanceLogRepository(db *sqlx.DB) domain.MaintenanceLogRepository {
	return &maintenanceLogRepository{db: db}
}

func (repo *maintenanceLogRepository) GetAll() (*[]domain.MAINTENANCE_LOG, error) {
	maintenanceLogs := make([]domain.MAINTENANCE_LOG, 0)
	err := repo.db.Select(&maintenanceLogs, "SELECT * FROM maintenance_log")
	if err != nil {
		return nil, err
	}
	return &maintenanceLogs, nil
}

func (repo *maintenanceLogRepository) GetByID(id int) (*domain.MAINTENANCE_LOG, error) {
	maintenanceLog := domain.MAINTENANCE_LOG{}
	err := repo.db.Get(&maintenanceLog, "SELECT * FROM maintenance_log WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &maintenanceLog, nil
}

func (repo *maintenanceLogRepository) GetByMaintenanceID(maintenance_id int) (*[]domain.MAINTENANCE_LOG, error) {
	maintenanceLogs := make([]domain.MAINTENANCE_LOG, 0)
	err := repo.db.Select(&maintenanceLogs, "SELECT * FROM maintenance_log WHERE maintenance_id = ?", maintenance_id)
	if err != nil {
		return nil, err
	}
	return &maintenanceLogs, nil
}

func (repo *maintenanceLogRepository) Create(maintenanceLog *domain.MAINTENANCE_LOG) error {
	t := repo.db.MustBegin()
	_, err := t.NamedExec("INSERT INTO maintenance_log (maintenance_id, staff_id, description, status) VALUES (:maintenance_id, :staff_id, :description, :status)", maintenanceLog)
	if err != nil {
		t.Rollback()
		return err
	}
	t.Commit()
	return nil
}

func (repo *maintenanceLogRepository) Update(maintenanceLog *domain.MAINTENANCE_LOG) error {
	t := repo.db.MustBegin()
	_, err := t.NamedExec("UPDATE maintenance_log SET maintenance_id = :maintenance_id, staff_id = :staff_id, description = :description, status = :status WHERE id = :id", maintenanceLog)
	if err != nil {
		t.Rollback()
		return err
	}
	t.Commit()
	return nil
}
