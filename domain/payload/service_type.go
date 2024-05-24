package payload

type ServiceTypeCreateDTO struct {
	NAME      string                        `json:"name" validate:"required"`
	DETAIL    string                        `json:"detail" validate:"required"`
	IS_ACTIVE bool                          `json:"isActive" validate:"boolean"`
	SERVICE   *[]ServiceOnCreateServiceType `json:"service" validate:"omitempty"`
}

type ServiceOnCreateServiceType struct {
	NAME        string  `json:"name" validate:"required"`
	DESCRIPTION string  `json:"description" validate:"required"`
	INFORMATION string  `json:"information" validate:"required"`
	PRICE       float32 `json:"price" validate:"required"`
	IS_ACTIVE   bool    `json:"isActive" validate:"boolean"`
}

type ServiceTypeUpdateDTO struct {
	ID        int    `json:"id" validate:"required"`
	NAME      string `json:"name" validate:"required"`
	DETAIL    string `json:"detail" validate:"required"`
	IS_ACTIVE bool   `json:"isActive" validate:"boolean"`
}

type ServiceTypeUpdateIsActiveDTO struct {
	ID        int  `json:"id" validate:"required"`
	IS_ACTIVE bool `json:"isActive" validate:"boolean"`
}
