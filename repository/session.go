package repository

import (
	"database/sql"
	"fmt"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/jmoiron/sqlx"
)

type sessionRepository struct {
	db *sqlx.DB
}

func NewSessionRepository(db *sqlx.DB) domain.SessionRepository {
	return &sessionRepository{db: db}
}

func (repo *sessionRepository) Get(ssid string) (*domain.Session, error) {
	var session domain.Session
	err := repo.db.Get(&session, "SELECT * FROM session WHERE id = ? LIMIT 1", ssid)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get session: %w", err)
	}
	return &session, nil
}

func (repo *sessionRepository) Create(session *domain.Session) error {
	_, err := repo.db.NamedExec("INSERT INTO session (id, user_id, ip_address, expired_at, created_at) VALUES (:id, :user_id,:ip_address, :expired_at, :created_at)", session)
	if err != nil {
		return fmt.Errorf("cannot query to create session: %w", err)
	}
	return nil
}

func (repo *sessionRepository) Delete(ssid string) error {
	_, err := repo.db.Exec("DELETE FROM session WHERE id = ?", ssid)
	if err != nil {
		return fmt.Errorf("cannot query to delete session: %w", err)
	}
	return nil
}
