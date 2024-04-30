package server

import (
	"log"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/config"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/constant"
	"github.com/SSSBoOm/CPE241_Project_Backend/server/controller"
	"github.com/SSSBoOm/CPE241_Project_Backend/server/middleware"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type FiberServer struct {
	app        *fiber.App
	cfg        *config.Config
	usecase    *domain.Usecase
	repository *domain.Repository
}

func NewFiberServer(
	cfg *config.Config,
	usecase *domain.Usecase,
	repository *domain.Repository,

) *FiberServer {
	return &FiberServer{
		cfg:        cfg,
		usecase:    usecase,
		repository: repository,
	}
}

func (s *FiberServer) Start() {
	app := fiber.New(fiber.Config{
		AppName: "API",
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	s.app = app
	s.Route()
	if err := app.Listen(":" + string(s.cfg.BACKEND_PORT)); err != nil {
		log.Fatal("Server is not running")
	}
}

func (s *FiberServer) Close() error {
	return s.app.Shutdown()
}

func (s *FiberServer) Route() {
	middlewareAuth := middleware.NewAuthMiddleware(s.usecase.SessionUsecase, s.usecase.UserUsecase, s.usecase.RoleUsecase)
	// AdminAuthMiddleware := middleware.NewRoleAuthMiddleware([]string{constant.ADMIN_ROLE})
	StaffAuthMiddleware := middleware.NewRoleAuthMiddleware([]string{constant.ADMIN_ROLE, constant.USER_ROLE})

	healthCheckController := controller.NewHealthCheckController()
	authController := controller.NewAuthController(s.usecase.AuthUsecase, s.usecase.GoogleUsecase, s.usecase.UserUsecase)
	userController := controller.NewUserController(s.usecase.UserUsecase)

	s.app.Use(swagger.New(swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path:     "/docs/swagger",
	}))
	api := s.app.Group("/api")

	api.Get("/healthcheck", healthCheckController.HealthCheck)

	auth := api.Group("/auth")
	auth.Get("/google", authController.GetUrl)
	auth.Get("/google/callback", authController.SignInWithGoogle)

	user := api.Group("/user")
	user.Get("/me", middlewareAuth, userController.Me)
	user.Get("/:id", middlewareAuth, StaffAuthMiddleware, userController.GetByID)
}
