package repository

import (
	"database/sql"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/jmoiron/sqlx"
)

type paymentTypeRepository struct {
	db *sqlx.DB
}

func NewPaymentTypeRepository(db *sqlx.DB) domain.PaymentTypeRepository {
	return &paymentTypeRepository{db: db}
}

func (repo *paymentTypeRepository) GetAll() (*[]domain.PaymentType, error) {
	paymentTypes := make([]domain.PaymentType, 0)
	err := repo.db.Select(&paymentTypes, "SELECT * FROM payment_type")
	if err != nil {
		return nil, err
	}
	return &paymentTypes, nil
}

func (repo *paymentTypeRepository) GetByID(id int) (*domain.PaymentType, error) {
	paymentType := domain.PaymentType{}
	err := repo.db.Get(&paymentType, "SELECT * FROM payment_type WHERE id = ?", id)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &paymentType, nil
}
