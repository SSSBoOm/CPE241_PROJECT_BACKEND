package domain

import "time"

type GenderType string

const (
	Male   GenderType = "Male"
	Female GenderType = "Female"
)

type User struct {
	Id        string     `json:"id" db:"id"`
	Email     string     `json:"email" db:"email"`
	Prefix    string     `json:"prefix" db:"prefix"`
	FirstName string     `json:"firstName" db:"first_name"`
	LastName  string     `json:"lastName" db:"last_name"`
	Gender    GenderType `json:"gender" db:"gender"`
	Phone     string     `json:"phone" db:"phone"`
	// Address    string    `json:"address" db:"address"`
	// City       string    `json:"city" db:"city"`
	// Postcode   string    `json:"postCode" db:"postcode"`
	// Country    string    `json:"country" db:"country"`
	ProfileUrl string    `json:"profileUrl" db:"profile_url"`
	CreatedAt  time.Time `json:"createdAt" db:"created_at"`
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
