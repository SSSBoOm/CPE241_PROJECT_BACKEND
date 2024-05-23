package domain

import "time"

type SERVICE struct {
	ID              int           `json:"id" db:"id"`
	NAME            string        `json:"name" db:"name"`
	DESCRIPTION     string        `json:"description" db:"description"`
	INFORMATION     string        `json:"information" db:"information"`
	PRICE           *float32      `json:"price" db:"price"`
	IS_ACTIVE       bool          `json:"isActive" db:"is_active"`
	SERVICE_TYPE_ID int           `json:"-" db:"service_type_id"`
	SERVICE_TYPE    *SERVICE_TYPE `json:"serviceType,omitempty" db:"-"`
	UPDATED_AT      time.Time     `json:"updateAt" db:"updated_at"`
	CREATED_AT      time.Time     `json:"createdAt" db:"created_at"`
}

type ServiceRepository interface {
	GetAll() (*[]SERVICE, error)
	GetById(id int) (*SERVICE, error)
	Create(service *SERVICE) error
	Update(service *SERVICE) error
	UpdateIsActive(id int, isActive bool) error
}

type ServiceUsecase interface {
	GetAll() (*[]SERVICE, error)
	GetById(id int) (*SERVICE, error)
	Create(service *SERVICE) error
	Update(service *SERVICE) error
	UpdateIsActive(id int, isActive bool) error
}
