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

func (repo *paymentTypeRepository) Create(paymentType *domain.PaymentType) error {
	_, err := repo.db.NamedExec("INSERT INTO payment_type (payment_type_name) VALUES (:payment_type_name)", paymentType)
	if err != nil {
		return err
	}
	return nil
}

func (repo *paymentTypeRepository) Update(paymentType *domain.PaymentType) error {
	_, err := repo.db.NamedExec("UPDATE payment_type SET payment_type_name = :payment_type_name WHERE id = :id", paymentType)
	if err != nil {
		return err
	}
	return nil
}
