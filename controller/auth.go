package controller

import (
	"fmt"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	env           *domain.Env
	authUsecase   domain.AuthUsecase
	googleUsecase domain.GoogleUsecase
	userUsecase   domain.UserUsecase
}

func NewAuthController(env *domain.Env, authUsecase domain.AuthUsecase, googleUsecase domain.GoogleUsecase, userUsecase domain.UserUsecase) *AuthController {
	return &AuthController{
		env:           env,
		authUsecase:   authUsecase,
		googleUsecase: googleUsecase,
		userUsecase:   userUsecase}
}

func (auth *AuthController) GetUrl(c *fiber.Ctx) error {
	path := auth.googleUsecase.GoogleConfig()
	url := path.AuthCodeURL("state")

	return c.Status(200).JSON(fiber.Map{"success": true, "url": url})
}

func (auth *AuthController) SignInWithGoogle(c *fiber.Ctx) error {
	_, _, err := auth.authUsecase.SignInWithGoogle(c)
	if err != nil {
		fmt.Println("SignIn ", err)
	}

	// c.Cookie(cookie)
	return c.Redirect("http://localhost:8080/")
}
