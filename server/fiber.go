package server

import (
	"log"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/config"
	"github.com/SSSBoOm/CPE241_Project_Backend/server/controller"
	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	app        *fiber.App
	cfg        *config.Config
	usecase    *domain.Usecase
	repository *domain.Repository
}

func NewFiberServer(
	usecase *domain.Usecase,
	repository *domain.Repository,

) *FiberServer {
	return &FiberServer{
		usecase:    usecase,
		repository: repository,
	}
}

func (s *FiberServer) Start() {
	app := fiber.New(fiber.Config{
		AppName: "API",
	})

	s.app = app

	s.Route()

	if err := app.Listen(":8080"); err != nil {
		log.Fatal("Server is not running")
	}
}

func (s *FiberServer) Close() error {
	return s.app.Shutdown()
}

func (s *FiberServer) Route() {
	authController := controller.NewAuthController(s.usecase.AuthUsecase, s.usecase.GoogleUsecase, s.usecase.UserUsecase)

	auth := s.app.Group("/api/auth")
	auth.Get("/google", authController.GetUrl)
	auth.Get("/google/callback", authController.SignInWithGoogle)
}
