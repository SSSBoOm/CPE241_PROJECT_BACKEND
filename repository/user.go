package repository

import (
	"database/sql"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindById(id string) (*domain.User, error) {
	user := domain.User{}
	err := r.db.Get(&user, "SELECT * FROM user WHERE id = ? LIMIT 1", id)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	user := domain.User{}
	err := r.db.Get(&user, "SELECT * FROM user WHERE email = ? LIMIT 1", email)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user *domain.User) error {
	_, err := r.db.NamedExec("INSERT INTO user (id, email, prefix, first_name, last_name, profile_url, phone, created_at) VALUES (:id, :email, :prefix, :first_name, :last_name, :profile_url, :phone, :created_at)", user)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) CreateFromGoogle(user *domain.User) error {
	_, err := r.db.NamedExec("INSERT INTO user (id, email, first_name, last_name, profile_url, created_at) VALUES (:id, :email, :first_name, :last_name, :profile_url, :created_at)", user)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Update(user *domain.User) error {
	_, err := r.db.NamedExec("UPDATE user SET prefix = :prefix, first_name = :first_name, last_name = :last_name, phone = :phone WHERE id = :id", user)
	if err != nil {
		return err
	}
	return nil
}
