package payload

type AddPaymentByUserIdDTO struct {
	ID              int    `json:"id" db:"id" validate:"required"`
	NAME            string `json:"name" db:"name" validate:"required"`
	PAYMENT_NAME    string `json:"paymentName" db:"payment_name" validate:"required"`
	PAYMENT_NUMBER  string `json:"paymentNumber" db:"payment_number" validate:"required"`
	PAYMENT_TYPE_ID int    `json:"-" db:"payment_type_id" validate:"required"`
}
