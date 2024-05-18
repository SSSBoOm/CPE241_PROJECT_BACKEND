package repository

import (
	"database/sql"
	"fmt"
	"time"

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
		return nil, fmt.Errorf("cannot query to get user: %w", err)
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	user := domain.User{}
	err := r.db.Get(&user, "SELECT * FROM user WHERE email = ? LIMIT 1", email)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get user: %w", err)
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
	user.PREFIX = ""
	user.ADDRESS = ""
	user.PHONE = ""
	user.GENDER = ""
	_, err := r.db.NamedExec("INSERT INTO user (id, email, first_name, last_name, profile_url, created_at, address, phone, gender) VALUES (:id, :email, :first_name, :last_name, :profile_url, :created_at, :address, :phone, :gender)", user)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) UpdateInfomation(user *domain.User) error {
	user.UPDATED_AT = time.Now()
	tx := r.db.MustBegin()
	_, err := tx.NamedExec("UPDATE user SET prefix = :prefix, first_name = :first_name, last_name = :last_name, gender = :gender, address = :address, dob = :dob, phone = :phone, updated_at = :updated_at WHERE id = :id", user)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r *userRepository) UpdateRoleById(userId string, roleID int) error {
	tx := r.db.MustBegin()
	_, err := tx.Exec("UPDATE user SET role_id = ?, updated_at = ? WHERE id = ?", roleID, time.Now(), userId)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
