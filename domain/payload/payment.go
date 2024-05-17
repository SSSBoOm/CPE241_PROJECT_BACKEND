package payload

type AddPaymentByUserIdDTO struct {
	ID              int    `json:"id" validate:"required"`
	NAME            string `json:"name" validate:"required"`
	PAYMENT_NAME    string `json:"paymentName" validate:"required"`
	PAYMENT_NUMBER  string `json:"paymentNumber" validate:"required"`
	PAYMENT_TYPE_ID int    `json:"-" validate:"required"`
}
