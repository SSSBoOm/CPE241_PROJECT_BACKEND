package repository

import (
	"time"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/jmoiron/sqlx"
)

type serviceTypeRepository struct {
	db *sqlx.DB
}

func NewServiceTypeRepository(db *sqlx.DB) domain.ServiceTypeRepository {
	return &serviceTypeRepository{db}
}

func (r *serviceTypeRepository) GetAll() (*[]domain.SERVICE_TYPE, error) {
	services := make([]domain.SERVICE_TYPE, 0)
	err := r.db.Select(&services, "SELECT * FROM service_type")
	if err != nil {
		return nil, err
	}

	return &services, nil
}

func (r *serviceTypeRepository) GetByID(id int) (*domain.SERVICE_TYPE, error) {
	var service domain.SERVICE_TYPE
	err := r.db.Get(&service, "SELECT * FROM service_type WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	return &service, nil
}

func (r *serviceTypeRepository) Create(serviceType *domain.SERVICE_TYPE) (*int, error) {
	t := r.db.MustBegin()
	data, err := t.NamedExec("INSERT INTO service_type (name, detail, is_active) VALUES (:name, :detail, :is_active)", serviceType)
	if err != nil {
		t.Rollback()
		return nil, err
	}
	t.Commit()
	rowId, _ := data.LastInsertId()
	Id := int(rowId)
	return &Id, nil
}

func (r *serviceTypeRepository) Update(serviceType *domain.SERVICE_TYPE) error {
	serviceType.UPDATED_AT = time.Now()
	t := r.db.MustBegin()
	_, err := t.NamedExec("UPDATE service_type SET name = :name, detail = :detail, is_active = :is_active, updated_at = :updated_at WHERE id = :id", serviceType)
	if err != nil {
		t.Rollback()
		return err
	}
	t.Commit()
	return nil
}

func (r *serviceTypeRepository) UpdateIsActive(id int, isActive bool) error {
	t := r.db.MustBegin()
	_, err := t.Exec("UPDATE service_type SET is_active = ?, updated_at = ? WHERE id = ?", isActive, time.Now(), id)
	if err != nil {
		t.Rollback()
		return err
	}
	t.Commit()
	return nil
}
