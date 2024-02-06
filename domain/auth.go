package domain

import "github.com/gofiber/fiber/v2"

type AuthUsecase interface {
	SignInWithGoogle(c *fiber.Ctx) (*User, *fiber.Cookie, error)
	// SignOut(sid string) (*fiber.Cookie, error)
}
