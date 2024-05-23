package payload

type AddPaymentByUserIdDTO struct {
	NAME            string `json:"name" validate:"required"`
	PAYMENT_FIRST_NAME    string `json:"paymentFirstName" validate:"required"`
	PAYMENT_LAST_NAME    string `json:"paymentLastName" validate:"required"`
	PAYMENT_NUMBER  string `json:"paymentNumber" validate:"required"`
	PAYMENT_TYPE_ID int    `json:"paymentTypeId" validate:"required"`
}
