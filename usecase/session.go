package usecase

import (
	"time"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/constant"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type sessionUsecase struct {
	sessionRepository domain.SessionRepository
}

func NewSessionUsecase(sessionRepository domain.SessionRepository) domain.SessionUsecase {
	return &sessionUsecase{
		sessionRepository: sessionRepository,
	}
}

func (u *sessionUsecase) Get(ssid string) (*domain.Session, error) {
	session, err := u.sessionRepository.Get(ssid)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (u *sessionUsecase) Create(userId string, ipAddress string) (*fiber.Cookie, error) {
	id := uuid.NewString()
	createdAt := time.Now()
	expiresAt := createdAt.Add(time.Hour * 24 * 7)

	if err := u.sessionRepository.Create(&domain.Session{
		ID:         id,
		USER_ID:    userId,
		IPADDRESS:  ipAddress,
		CREATED_AT: createdAt,
		EXPIRED_AT: expiresAt,
	}); err != nil {
		return nil, err
	}

	cookie := &fiber.Cookie{
		Name:     constant.SessionCookieName,
		Value:    id,
		HTTPOnly: true,
		Secure:   true,
		Expires:  expiresAt,
	}
	return cookie, nil
}

func (u *sessionUsecase) Delete(ssid string) error {
	if err := u.sessionRepository.Delete(ssid); err != nil {
		return err
	}
	return nil
}
