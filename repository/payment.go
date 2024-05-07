package repository

import (
	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/jmoiron/sqlx"
)

type paymentRepository struct {
	db *sqlx.DB
}

func NewPaymentRepository(db *sqlx.DB) domain.PaymentRepository {
	return &paymentRepository{db: db}
}

func (repo *paymentRepository) GetAll() (*[]domain.Payment, error) {
	payments := make([]domain.Payment, 0)
	err := repo.db.Select(&payments, "SELECT * FROM payment")
	if err != nil {
		return nil, err
	}
	return &payments, nil
}

func (repo *paymentRepository) GetByID(id int) (*domain.Payment, error) {
	payment := domain.Payment{}
	err := repo.db.Get(&payment, "SELECT * FROM payment WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (repo *paymentRepository) Create(payment *domain.Payment) error {
	_, err := repo.db.NamedExec("INSERT INTO payment (name, payment_name, payment_number, user_id, payment_type_id) VALUES (:name, :payment_name, :payment_number, :user_id, :payment_type_id)", payment)
	if err != nil {
		return err
	}
	return nil
}
