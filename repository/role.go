package repository

import (
	"database/sql"
	"fmt"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/jmoiron/sqlx"
)

type roleRepository struct {
	db *sqlx.DB
}

func NewRoleRepository(db *sqlx.DB) domain.RoleRepository {
	return &roleRepository{db: db}
}

func (repo *roleRepository) Get(id int) (*domain.Role, error) {
	var role domain.Role
	err := repo.db.Get(&role, "SELECT * FROM role WHERE id = ? LIMIT 1", id)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get role: %w", err)
	}
	return &role, nil
}
