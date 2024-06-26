package domain

import "time"

type SERVICE_TYPE struct {
	ID         int        `json:"id" db:"id"`
	NAME       string     `json:"name" db:"name"`
	DETAIL     string     `json:"detail" db:"detail"`
	SERVICE    *[]SERVICE `json:"service" db:"-"`
	IS_ACTIVE  bool       `json:"isActive" db:"is_active"`
	UPDATED_AT time.Time  `json:"updateAt" db:"updated_at"`
	CREATED_AT time.Time  `json:"createdAt" db:"created_at"`
}

type ServiceTypeRepository interface {
	GetAll() (*[]SERVICE_TYPE, error)
	GetByID(id int) (*SERVICE_TYPE, error)
	Create(roomType *SERVICE_TYPE) (*int, error)
	Update(roomType *SERVICE_TYPE) error
	UpdateIsActive(id int, isActive bool) error
}

type ServiceTypeUsecase interface {
	GetAll() (*[]SERVICE_TYPE, error)
	GetByID(id int) (*SERVICE_TYPE, error)
	Create(roomType *SERVICE_TYPE) (*int, error)
	Update(roomType *SERVICE_TYPE) error
	UpdateIsActive(id int, isActive bool) error
}
