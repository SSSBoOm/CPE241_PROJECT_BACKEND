package domain

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type SESSION_TYPE string

const (
	USER  SESSION_TYPE = "USER"
	STAFF SESSION_TYPE = "STAFF"
)

type Session struct {
	ID         string       `json:"id" db:"id"`
	USER_ID    string       `json:"userId" db:"user_id"`
	TYPE       SESSION_TYPE `json:"-" db:"session_type"`
	IPADDRESS  string       `json:"-" db:"ip_address"`
	CREATED_AT time.Time    `json:"createdAt" db:"created_at"`
	EXPIRED_AT time.Time    `json:"expiredAt" db:"expired_at"`
}

type SessionRepository interface {
	Create(session *Session) error
	Get(ssid string) (*Session, error)
}

type SessionUsecase interface {
	Create(userId string, ipAddress string, sessionType SESSION_TYPE) (*fiber.Cookie, error)
}
