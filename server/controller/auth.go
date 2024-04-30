package controller

import (
	"fmt"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authUsecase   domain.AuthUsecase
	googleUsecase domain.GoogleUsecase
	userUsecase   domain.UserUsecase
}

func NewAuthController(authUsecase domain.AuthUsecase, googleUsecase domain.GoogleUsecase, userUsecase domain.UserUsecase) *AuthController {
	return &AuthController{
		authUsecase:   authUsecase,
		googleUsecase: googleUsecase,
		userUsecase:   userUsecase}
}

func (auth *AuthController) GetUrl(c *fiber.Ctx) error {
	path := auth.googleUsecase.GoogleConfig()
	url := path.AuthCodeURL("state")

	return c.Status(200).JSON(fiber.Map{"success": true, "url": url})
}

func (auth *AuthController) SignInWithGoogle(ctx *fiber.Ctx) error {
	cookie, err := auth.authUsecase.SignInWithGoogle(ctx)
	if err != nil {
		fmt.Println("SignIn :", err)
	}

	ctx.Cookie(cookie)
	return ctx.Redirect("http://localhost:8080/")
}
