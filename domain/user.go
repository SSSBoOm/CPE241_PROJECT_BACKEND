package domain

import (
	"time"
)

type PrefixType string

const (
	MR      PrefixType = "MR"
	MRS     PrefixType = "MRS"
	MS      PrefixType = "MS"
	DEFAULT PrefixType = ""
)

type GenderType string

const (
	MALE   GenderType = "MALE"
	FEMALE GenderType = "FEMALE"
)

type User struct {
	ID          string     `json:"id,omitempty" db:"id"`
	EMAIL       string     `json:"email" db:"email"`
	PREFIX      PrefixType `json:"prefix" db:"prefix"`
	FIRST_NAME  string     `json:"firstName" db:"first_name"`
	LAST_NAME   string     `json:"lastName" db:"last_name"`
	DOB         *time.Time `json:"dob" db:"dob"`
	PHONE       string     `json:"phone" db:"phone"`
	GENDER      GenderType `json:"gender" db:"gender"`
	ADDRESS     string     `json:"address" db:"address"`
	PROFILE_URL string     `json:"profileUrl" db:"profile_url"`
	ROLE_ID     int        `json:"-" db:"role_id"`
	UPDATED_AT  time.Time  `json:"updateAt" db:"updated_at"`
	CREATED_AT  time.Time  `json:"createdAt" db:"created_at"`
}

type UserRepository interface {
	GetAll() (*[]User, error)
	FindById(id string) (*User, error)
	FindByEmail(email string) (*User, error)
	Create(user *User) error
	CreateFromGoogle(user *User) error
	UpdateInfomation(user *User) error
	UpdateRoleById(id string, roleID int) error
}

type UserUsecase interface {
	GetAll() (*[]User, error)
	CreateFromGoogle(name string, email string, picture string) (*User, error)
	FindById(id string) (*User, error)
	FindByEmail(email string) (*User, error)
	UpdateRoleById(id string, roleID int) error
	UpdateInfomation(user *User) error
}
