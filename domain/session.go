package domain

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Session struct {
	ID         string    `json:"id" db:"id"`
	USER_ID    string    `json:"userId" db:"user_id"`
	IPADDRESS  string    `json:"-" db:"ip_address"`
	CREATED_AT time.Time `json:"createdAt" db:"created_at"`
	EXPIRED_AT time.Time `json:"expiredAt" db:"expired_at"`
}

type SessionRepository interface {
	Create(session *Session) error
	Get(ssid string) (*Session, error)
}

type SessionUsecase interface {
	Create(userId string, ipAddress string) (*fiber.Cookie, error)
	// Get(ssid string) (*Session, error)
}
