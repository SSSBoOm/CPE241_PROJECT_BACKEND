package domain

type PaymentType struct {
	ID   int    `json:"id" db:"id"`
	NAME string `json:"name" db:"payment_type_name"`
}

type PaymentTypeRepository interface {
	GetAll() (*[]PaymentType, error)
	GetByID(id int) (*PaymentType, error)
	Create(paymentType *PaymentType) error
	Update(paymentType *PaymentType) error
}

type PaymentTypeUsecase interface {
	GetAll() (*[]PaymentType, error)
	GetByID(id int) (*PaymentType, error)
	Create(paymentType *PaymentType) error
	Update(paymentType *PaymentType) error
}
