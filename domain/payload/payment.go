package payload

type AddPaymentByUserIdDTO struct {
	NAME            string `json:"name" validate:"required"`
	PAYMENT_NAME    string `json:"paymentName" validate:"required"`
	PAYMENT_NUMBER  string `json:"paymentNumber" validate:"required"`
	PAYMENT_TYPE_ID int    `json:"payment_type_id" validate:"required"`
}
