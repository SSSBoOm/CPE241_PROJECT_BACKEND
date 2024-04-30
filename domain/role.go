package domain

type Role struct {
	ID   int    `json:"-" db:"id"`
	NAME string `json:"-" db:"name"`
}

type RoleRepository interface {
	Get(id int) (*Role, error)
}

type RoleUsecase interface {
	Get(id int) (*Role, error)
}
