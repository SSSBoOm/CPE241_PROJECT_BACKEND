package payload

type ServiceTypeCreateDTO struct {
	NAME      string  `json:"name" validate:"required"`
	DETAIL    string  `json:"detail" validate:"required"`
	IS_ACTIVE bool    `json:"isActive" validate:"required"`
}

type ServiceTypeUpdateDTO struct {
	ID        int     `json:"id" validate:"required"`
	NAME      string  `json:"name" validate:"required"`
	DETAIL    string  `json:"detail" validate:"required"`
	IS_ACTIVE bool    `json:"isActive" validate:"boolean"`
}

type ServiceTypeUpdateIsActiveDTO struct {
	ID        int  `json:"id" validate:"required"`
	IS_ACTIVE bool `json:"isActive" validate:"boolean"`
}