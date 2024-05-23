package repository

import (
	"time"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/jmoiron/sqlx"
)

type serviceRepository struct {
	db *sqlx.DB
}

func NewServiceRepository(db *sqlx.DB) domain.ServiceRepository {
	return &serviceRepository{db}
}

func (r *serviceRepository) GetAll() (*[]domain.SERVICE, error) {
	services := make([]domain.SERVICE, 0)
	err := r.db.Select(&services, "SELECT * FROM service")
	if err != nil {
		return nil, err
	}

	return &services, nil
}

func (r *serviceRepository) GetById(id int) (*domain.SERVICE, error) {
	var service domain.SERVICE
	err := r.db.Get(&service, "SELECT * FROM service WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	return &service, nil
}

func (r *serviceRepository) Create(service *domain.SERVICE) error {
	t := r.db.MustBegin()
	_, err := t.NamedExec("INSERT INTO service (name, description, information, is_active, price, service_type_id) VALUES (:name, :description, :information, :price, :is_active, :service_type_id)", service)
	if err != nil {
		t.Rollback()
		return err
	}
	t.Commit()
	return nil
}

func (r *serviceRepository) Update(service *domain.SERVICE) error {
	service.UPDATED_AT = time.Now()
	t := r.db.MustBegin()
	_, err := t.NamedExec("UPDATE service SET name = :name, description = :description, information = :information, price = :price, is_active = :is_active, service_type_id = :service_type_id, updated_at = :updated_at WHERE id = :id", service)
	if err != nil {
		t.Rollback()
		return err
	}
	t.Commit()
	return nil
}

func (r *serviceRepository) UpdateIsActive(id int, isActive bool) error {
	t := r.db.MustBegin()
	_, err := t.Exec("UPDATE service SET is_active = ?, updated_at = ? WHERE id = ?", isActive, time.Now(), id)
	if err != nil {
		t.Rollback()
		return err
	}
	t.Commit()
	return nil
}
