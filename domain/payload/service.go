package payload

type ServiceCreateDTO struct {
	NAME            string  `json:"name" validate:"required"`
	DESCRIPTION     string  `json:"description" validate:"required"`
	INFORMATION     string  `json:"information" validate:"required"`
	PRICE           float32 `json:"price" validate:"required"`
	IMAGE_URL       string  `json:"imageUrl" validate:"required"`
	IS_ACTIVE       bool    `json:"isActive" validate:"boolean"`
	SERVICE_TYPE_ID int     `json:"serviceTypeId" validate:"required"`
}

type ServiceUpdateDTO struct {
	ID              int     `json:"id" validate:"required"`
	NAME            string  `json:"name" validate:"required"`
	DESCRIPTION     string  `json:"description" validate:"required"`
	INFORMATION     string  `json:"information" validate:"required"`
	IMAGE_URL       string  `json:"imageUrl" validate:"required"`
	PRICE           float32 `json:"price" validate:"required"`
	SERVICE_TYPE_ID int     `json:"serviceTypeId" validate:"required"`
	IS_ACTIVE       bool    `json:"isActive" validate:"boolean"`
}

type ServiceUpdateIsActiveDTO struct {
	ID        int  `json:"id" validate:"required"`
	IS_ACTIVE bool `json:"isActive" validate:"boolean"`
}
