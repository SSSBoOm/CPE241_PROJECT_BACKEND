package domain

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Session struct {
	Id        string    `json:"id" db:"id"`
	UserId    string    `json:"userId" db:"user_id"`
	IpAddress string    `json:"-" db:"ip_address"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	ExpiredAt time.Time `json:"expiredAt" db:"expired_at"`
}

type SessionRepository interface {
	Create(session *Session) error
	Get(ssid string) (*Session, error)
}

type SessionUsecase interface {
	Create(userId string, ipAddress string) (*fiber.Cookie, error)
}
