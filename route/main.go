package route

import (
	"github.com/SSSBoOm/CPE241_Project_Backend/controller"
	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/SSSBoOm/CPE241_Project_Backend/repository"
	"github.com/SSSBoOm/CPE241_Project_Backend/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func MainRoute(app *fiber.App, db *sqlx.DB, env *domain.Env) {
	userRepository := repository.NewUserRepository(db)

	googleUsecase := usecase.NewGoogleUsecase(env)
	userUsecase := usecase.NewUserUsecase(userRepository)
	authUsecase := usecase.NewAuthUsecase(googleUsecase, userUsecase)

	authController := controller.NewAuthController(env, authUsecase, googleUsecase, userUsecase)

	auth := app.Group("/api/auth")
	auth.Get("/google", authController.GetUrl)
	auth.Get("/google/callback", authController.SignInWithGoogle)
}
