package payload

import (
	"time"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
)

type UpdateUserDTO struct {
	ID         string            `json:"id" validate:"required"`
	PREFIX     string            `json:"prefix" validate:"required"`
	FIRST_NAME string            `json:"firstName" validate:"required"`
	LAST_NAME  string            `json:"lastName" validate:"required"`
	DOB        time.Time         `json:"dob" validate:"required"`
	PHONE      string            `json:"phone" validate:"min=10"`
	GENDER     domain.GenderType `json:"gender" validate:"required"`
	ADDRESS    string            `json:"address" validate:"required"`
}

type UpdateUserInformationDTO struct {
	PREFIX     string            `json:"prefix" validate:"required"`
	FIRST_NAME string            `json:"firstName" validate:"required"`
	LAST_NAME  string            `json:"lastName" validate:"required"`
	DOB        time.Time         `json:"dob" validate:"required"`
	PHONE      string            `json:"phone" validate:"min=10"`
	GENDER     domain.GenderType `json:"gender" validate:"required"`
	ADDRESS    string            `json:"address" validate:"required"`
}

type UpdateUserRoleDTO struct {
	USER_ID string `json:"userId" validate:"required"`
	ROLE_ID int    `json:"roleId" validate:"required"`
}
