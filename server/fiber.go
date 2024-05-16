package server

import (
	"log"

	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/config"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/constant"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/validator"
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
	validator := validator.NewPayloadValidator()

	middlewareAuth := middleware.NewAuthMiddleware(s.usecase.SessionUsecase, s.usecase.UserUsecase, s.usecase.RoleUsecase)
	AdminAuthMiddleware := middleware.NewRoleAuthMiddleware([]string{constant.ADMIN_ROLE})
	StaffAuthMiddleware := middleware.NewRoleAuthMiddleware([]string{constant.ADMIN_ROLE, constant.USER_ROLE})

	healthCheckController := controller.NewHealthCheckController()
	authController := controller.NewAuthController(s.cfg, s.usecase.AuthUsecase, s.usecase.GoogleUsecase, s.usecase.UserUsecase)
	userController := controller.NewUserController(validator, s.usecase.UserUsecase, s.usecase.PaymentUsecase)
	roleController := controller.NewRoleController(s.usecase.RoleUsecase)
	paymentController := controller.NewPaymentController(s.usecase.PaymentUsecase)
	paymentTypeController := controller.NewPaymentTypeController(s.usecase.PaymentTypeUsecase)
	roomTypeController := controller.NewRoomTypeController(validator, s.usecase.RoomTypeUsecase)
	roomController := controller.NewRoomController(validator, s.usecase.RoomUsecase)
	maintenanceController := controller.NewMaintenanceController(validator, s.usecase.MaintenanceUsecase)

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
	auth.Post("/logout", authController.SignOut)

	user := api.Group("/user")
	user.Get("/me", middlewareAuth, userController.Me)
	user.Get("/payment", middlewareAuth, userController.GetPaymentByUserID)
	user.Patch("/", middlewareAuth, userController.UpdateInfomationByID)

	role := api.Group("/role")
	role.Get("/all", middlewareAuth, AdminAuthMiddleware, roleController.GetALL)

	admin := api.Group("/admin")
	admin.Get("/manage/user/:id", middlewareAuth, StaffAuthMiddleware, userController.GetByID)
	admin.Put("/manage/user", middlewareAuth, AdminAuthMiddleware, userController.UpdateByID)
	admin.Put("/manage/role", middlewareAuth, AdminAuthMiddleware, userController.UpdateRoleByID)

	payment := api.Group("/payment")
	payment.Post("/", middlewareAuth, paymentController.AddPaymentByUser)

	paymentType := api.Group("/payment_type")
	paymentType.Get("/all", middlewareAuth, paymentTypeController.GetAll)

	room := api.Group("/room")
	room.Get("/all", middlewareAuth, StaffAuthMiddleware, roomTypeController.GetRoomTypeList)
	room.Get("/:id", middlewareAuth, StaffAuthMiddleware, roomTypeController.GetRoomTypeByID)
	room.Post("/", middlewareAuth, AdminAuthMiddleware, roomController.Create)
	room.Put("/:id", middlewareAuth, AdminAuthMiddleware, roomTypeController.UpdateRoomType)

	roomType := api.Group("/room_type")
	roomType.Get("/all", roomTypeController.GetRoomTypeList)
	roomType.Get("/:id", roomTypeController.GetRoomTypeByID)
	roomType.Post("/", middlewareAuth, AdminAuthMiddleware, roomTypeController.CreateRoomType)
	roomType.Put("/:id", middlewareAuth, AdminAuthMiddleware, roomTypeController.UpdateRoomType)

	maintenance := api.Group("/maintenance")
	maintenance.Post("/", middlewareAuth, StaffAuthMiddleware, maintenanceController.Create)
}
