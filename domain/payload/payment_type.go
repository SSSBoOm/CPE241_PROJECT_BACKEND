package payload

type PaymentTypeCreateDTO struct {
	NAME string `json:"name" validate:"required"`
}

type PaymentTypeUpdateDTO struct {
	ID   int    `json:"id" validate:"required"`
	NAME string `json:"name" validate:"required"`
}
