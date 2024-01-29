package domain

import "time"

type User struct {
	Id         string    `json:"id" db:"id"`
	Email      string    `json:"email" db:"email"`
	Prefix     string    `json:"prefix" db:"prefix"`
	FirstName  string    `json:"firstName" db:"first_name"`
	LastName   string    `json:"lastName" db:"last_name"`
	ProfileUrl string    `json:"profileUrl" db:"profile_url"`
	Phone      string    `json:"phone" db:"phone"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

type UserRepository interface {
	FindById(id string) (*User, error)
	FindByEmail(email string) (*User, error)
	Create(user *User) error
	Update(user *User) error
}

type UserUsecase interface {
	CreateFromGoogle(profile *GoogleResponse) (*User, error)
	FindById(id string) (*User, error)
	FindByEmail(email string) (*User, error)
}
