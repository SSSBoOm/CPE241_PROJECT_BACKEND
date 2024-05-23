package repository

import (
	"time"

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
	err := repo.db.Get(&payment, "SELECT * FROM payment WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (repo *paymentRepository) GetByUserID(userId string) (*[]domain.Payment, error) {
	payments := make([]domain.Payment, 0)
	err := repo.db.Select(&payments, "SELECT * FROM payment WHERE user_id = ?", userId)
	if err != nil {
		return nil, err
	}
	return &payments, nil
}

func (repo *paymentRepository) Create(payment *domain.Payment) error {
	t := repo.db.MustBegin()
	_, err := t.NamedExec("INSERT INTO payment (name, payment_first_name, payment_last_name, payment_number, user_id, payment_type_id) VALUES (:name, :payment_first_name, :payment_last_name, :payment_number, :user_id, :payment_type_id)", payment)
	if err != nil {
		t.Rollback()
		return err
	}
	t.Commit()
	return nil
}

func (repo *paymentRepository) Update(payment *domain.Payment) error {
	payment.UPDATED_AT = time.Now()
	t := repo.db.MustBegin()
	_, err := t.NamedExec("UPDATE payment SET name = :name, payment_first_name = :payment_first_name, payment_last_name = :payment_last_name, payment_number = :payment_number, user_id = :user_id, payment_type_id = :payment_type_id, updated_at := updated_at WHERE id = :id", payment)
	if err != nil {
		t.Rollback()
		return err
	}
	t.Commit()
	return nil
}
