package domain

import "time"

type GenderType string

const (
	MALE   GenderType = "MALE"
	FEMALE GenderType = "FEMALE"
)

type User struct {
	ID          string     `json:"id" db:"id"`
	EMAIL       string     `json:"email" db:"email"`
	PREFIX      string     `json:"prefix" db:"prefix"`
	FIRST_NAME  string     `json:"firstName" db:"first_name"`
	LAST_NAME   string     `json:"lastName" db:"last_name"`
	DOB         time.Time  `json:"dob" db:"dob"`
	PHONE       string     `json:"phone" db:"phone"`
	GENDER      GenderType `json:"gender" db:"gender"`
	PROFILE_URL string     `json:"profileUrl" db:"profile_url"`
	CREATE_AT   time.Time  `json:"createdAt" db:"created_at"`
}

type UserRepository interface {
	FindById(id string) (*User, error)
	FindByEmail(email string) (*User, error)
	Create(user *User) error
	CreateFromGoogle(user *User) error
	Update(user *User) error
}

type UserUsecase interface {
	CreateFromGoogle(name string, email string, picture string) (*User, error)
	FindById(id string) (*User, error)
	FindByEmail(email string) (*User, error)
}
