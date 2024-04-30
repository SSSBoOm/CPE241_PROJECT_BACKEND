package payload

import (
	"time"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
)

type UpdateUserDTO struct {
	PREFIX     string            `json:"prefix" db:"prefix" validate:"required"`
	FIRST_NAME string            `json:"firstName" db:"first_name" validate:"required"`
	LAST_NAME  string            `json:"lastName" db:"last_name" validate:"required"`
	DOB        time.Time         `json:"dob" db:"dob" validate:"required"`
	PHONE      string            `json:"phone" db:"phone" validate:"min=10"`
	GENDER     domain.GenderType `json:"gender" db:"gender" validate:"required"`
	ADDRESS    string            `json:"address" db:"address" validate:"required"`
}
