package payload

import (
	"time"
)

type AddPaymentByUserIdDTO struct {
	ID              int       `json:"id" db:"id"`
	NAME            string    `json:"name" db:"name"`
	PAYMENT_NAME    string    `json:"paymentName" db:"payment_name"`
	PAYMENT_NUMBER  string    `json:"paymentNumber" db:"payment_number"`
	PAYMENT_TYPE_ID int       `json:"-" db:"payment_type_id"`
	CREATE_AT       time.Time `json:"createdAt" db:"created_at"`
}
