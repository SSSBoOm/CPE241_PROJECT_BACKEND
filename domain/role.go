package domain

type Role struct {
	ID   int    `json:"id" db:"id"`
	NAME string `json:"name" db:"name"`
}

type RoleRepository interface {
	Get(id int) (*Role, error)
	GetAll() (*[]Role, error)
}

type RoleUsecase interface {
	Get(id int) (*Role, error)
	GetAll() (*[]Role, error)
}
