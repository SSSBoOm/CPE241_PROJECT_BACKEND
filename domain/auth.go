package domain

import "github.com/gofiber/fiber/v2"

type AuthUsecase interface {
	GoogleLogin(c *fiber.Ctx) (*User, *fiber.Cookie, error)
	// Logout(sid string) (*fiber.Cookie, error)
}
