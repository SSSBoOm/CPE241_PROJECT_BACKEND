package domain

import "time"

type Payment struct {
	ID                 int         `json:"id" db:"id"`
	NAME               string      `json:"name" db:"name"`
	PAYMENT_FIRST_NAME string      `json:"paymentFirstName" db:"payment_first_name"`
	PAYMENT_LAST_NAME  string      `json:"paymentLastName" db:"payment_last_name"`
	PAYMENT_NUMBER     string      `json:"paymentNumber" db:"payment_number"`
	USER_ID            string      `json:"userId" db:"user_id"`
	PAYMENT_TYPE_ID    int         `json:"-" db:"payment_type_id"`
	PAYMENT_TYPE       PaymentType `json:"paymentType"`
	UPDATED_AT         time.Time   `json:"updatedAt" db:"updated_at"`
	CREATED_AT         time.Time   `json:"createdAt" db:"created_at"`
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
