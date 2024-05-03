package controller

import (
	"fmt"
	"time"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/config"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/constant"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	config         *config.Config
	authUsecase    domain.AuthUsecase
	googleUsecase  domain.GoogleUsecase
	userUsecase    domain.UserUsecase
	sessionUsecase domain.SessionUsecase
}

func NewAuthController(config *config.Config, authUsecase domain.AuthUsecase, googleUsecase domain.GoogleUsecase, userUsecase domain.UserUsecase) *AuthController {
	return &AuthController{
		config:        config,
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
	return ctx.Redirect(auth.config.FRONTEND_URL)
}

func (auth *AuthController) SignOut(ctx *fiber.Ctx) error {
	ssid := ctx.Cookies(constant.SessionCookieName)
	if err := auth.sessionUsecase.Delete(ssid); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(domain.Response{
			SUCCESS: false,
			MESSAGE: constant.MESSAGE_INTERNAL_SERVER_ERROR,
		})
	}

	ctx.Cookie(&fiber.Cookie{Name: constant.SessionCookieName, Expires: time.Unix(0, 0)})
	return ctx.Status(fiber.StatusOK).Redirect(auth.config.FRONTEND_URL)
}
