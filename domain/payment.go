package domain

import "time"

type Payment struct {
	ID              int         `json:"id" db:"id"`
	NAME            string      `json:"name" db:"name"`
	PAYMENT_NAME    string      `json:"paymentName" db:"payment_name"`
	PAYMENT_NUMBER  string      `json:"paymentNumber" db:"payment_number"`
	USER_ID         string      `json:"userId" db:"user_id"`
	PAYMENT_TYPE_ID int         `json:"-" db:"payment_type_id"`
	PAYMENT_TYPE    PaymentType `json:"paymentType"`
	CREATED_AT      time.Time   `json:"createdAt" db:"created_at"`
}

type PaymentRepository interface {
	GetAll() (*[]Payment, error)
	GetByID(id int) (*Payment, error)
	GetByUserID(userId string) (*[]Payment, error)
	Create(payment *Payment) error
}

type PaymentUsecase interface {
	GetAll() (*[]Payment, error)
	GetByID(id int) (*Payment, error)
	Create(payment *Payment) error
	GetByUserID(userId string) (*[]Payment, error)
}
