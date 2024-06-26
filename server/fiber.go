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
		AllowOrigins:     "https://hotel.sssboom.xyz/,http://localhost:5173",
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowCredentials: true,
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
	authController := controller.NewAuthController(s.cfg, s.usecase.AuthUsecase, s.usecase.GoogleUsecase, s.usecase.UserUsecase, s.usecase.SessionUsecase)
	userController := controller.NewUserController(validator, s.usecase.UserUsecase, s.usecase.PaymentUsecase, s.usecase.ReservationUsecase)
	roleController := controller.NewRoleController(s.usecase.RoleUsecase)
	paymentController := controller.NewPaymentController(validator, s.usecase.PaymentUsecase)
	paymentTypeController := controller.NewPaymentTypeController(validator, s.usecase.PaymentTypeUsecase)
	roomTypeController := controller.NewRoomTypeController(validator, s.usecase.RoomTypeUsecase)
	roomController := controller.NewRoomController(validator, s.usecase.RoomUsecase)
	maintenanceController := controller.NewMaintenanceController(validator, s.usecase.MaintenanceUsecase)
	maintenanceLogController := controller.NewMaintenanceLogController(validator, s.usecase.MaintenanceUsecase, s.usecase.MaintenanceLogUsecase)
	reservationController := controller.NewReservationController(validator, s.usecase.ReservationUsecase, s.usecase.RoomUsecase, s.usecase.RoomTypeUsecase)
	reservationTaskController := controller.NewReservationTaskController(validator, s.usecase.ReservationTaskUsecase)
	serviceTypeController := controller.NewServiceTypeController(validator, s.usecase.ServiceTypeUsecase)
	serviceController := controller.NewServiceController(validator, s.usecase.ServiceUsecase, s.usecase.ServiceTypeUsecase)
	promotionPriceController := controller.NewPromotionPriceController(validator, s.usecase.PromotionPriceUsecase, s.usecase.RoomTypePromotionPriceUsecase)
	roomTypePromotionPriceController := controller.NewRoomTypePromotionPriceController(validator, s.usecase.RoomTypePromotionPriceUsecase)
	dashboardController := controller.NewDashboardController(validator, s.usecase.DashboardUsecase)

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
	auth.Get("/logout", middlewareAuth, authController.SignOut)

	user := api.Group("/user")
	user.Get("/me", middlewareAuth, userController.Me)
	user.Get("/payment", middlewareAuth, userController.GetPaymentByUserID)
	user.Patch("/", middlewareAuth, userController.UpdateInfomationByID)

	role := api.Group("/role")
	role.Get("/", middlewareAuth, AdminAuthMiddleware, roleController.GetALL)

	admin := api.Group("/admin")
	admin.Get("/manage/user", middlewareAuth, StaffAuthMiddleware, userController.GetALL)
	admin.Get("/manage/user", middlewareAuth, AdminAuthMiddleware, userController.UpdateByID)
	admin.Get("/manage/user/:id", middlewareAuth, StaffAuthMiddleware, userController.GetByID)
	admin.Put("/manage/user", middlewareAuth, AdminAuthMiddleware, userController.UpdateByID)
	admin.Put("/manage/role", middlewareAuth, AdminAuthMiddleware, userController.UpdateRoleByID)

	payment := api.Group("/payment")
	payment.Post("/", middlewareAuth, paymentController.AddPaymentByUser)

	paymentType := api.Group("/payment_type")
	paymentType.Get("/", paymentTypeController.GetAll)
	paymentType.Get("/:id", middlewareAuth, paymentTypeController.GetByID)
	paymentType.Post("/", middlewareAuth, AdminAuthMiddleware, paymentTypeController.Create)
	paymentType.Put("/:id", middlewareAuth, AdminAuthMiddleware, paymentTypeController.Update)

	room := api.Group("/room")
	room.Get("/", middlewareAuth, StaffAuthMiddleware, roomController.GetAll)
	room.Get("/:id", middlewareAuth, StaffAuthMiddleware, roomController.GetByID)
	room.Post("/", middlewareAuth, AdminAuthMiddleware, roomController.Create)
	room.Post("/active", middlewareAuth, AdminAuthMiddleware, roomController.UpdateIsActive)
	room.Put("/:id", middlewareAuth, AdminAuthMiddleware, roomController.Update)

	roomType := api.Group("/room_type")
	roomType.Get("/", roomTypeController.GetRoomTypeList)
	roomType.Get("/:id", roomTypeController.GetRoomTypeByID)
	roomType.Post("/", middlewareAuth, AdminAuthMiddleware, roomTypeController.CreateRoomType)
	roomType.Post("/active", middlewareAuth, AdminAuthMiddleware, roomTypeController.UpdateRoomTypeIsActive)
	roomType.Put("/:id", middlewareAuth, AdminAuthMiddleware, roomTypeController.UpdateRoomType)

	serviceType := api.Group("/service_type")
	serviceType.Get("/", serviceTypeController.GetAll)
	serviceType.Get("/:id", middlewareAuth, StaffAuthMiddleware, serviceTypeController.GetByID)
	serviceType.Post("/", middlewareAuth, AdminAuthMiddleware, serviceTypeController.Create)
	serviceType.Post("/active", middlewareAuth, AdminAuthMiddleware, serviceTypeController.UpdateIsActive)
	serviceType.Put("/:id", middlewareAuth, AdminAuthMiddleware, serviceTypeController.Update)

	service := api.Group("/service")
	service.Get("/", serviceController.GetAll)
	service.Get("/:id", middlewareAuth, StaffAuthMiddleware, serviceController.GetByID)
	service.Post("/", middlewareAuth, AdminAuthMiddleware, serviceController.Create)
	service.Post("/active", middlewareAuth, AdminAuthMiddleware, serviceController.UpdateIsActive)
	service.Put("/:id", middlewareAuth, AdminAuthMiddleware, serviceController.Update)

	maintenance := api.Group("/maintenance")
	maintenance.Get("/", middlewareAuth, StaffAuthMiddleware, maintenanceController.GetAll)
	maintenance.Get("/:id", middlewareAuth, StaffAuthMiddleware, maintenanceController.GetByID)
	maintenance.Post("/", middlewareAuth, StaffAuthMiddleware, maintenanceController.Create)

	maintenanceLog := api.Group("/maintenance_log")
	maintenanceLog.Post("/", middlewareAuth, StaffAuthMiddleware, maintenanceLogController.Create)

	reservation := api.Group("/reservation")
	reservation.Get("/", middlewareAuth, StaffAuthMiddleware, reservationController.GetAll)
	reservation.Get("/type/:type", reservationController.GetReservationByReservationType)
	reservation.Get("/me", middlewareAuth, StaffAuthMiddleware, reservationController.GetReservationByUserID)
	reservation.Get("/:id", middlewareAuth, StaffAuthMiddleware, reservationController.GetByID)
	reservation.Post("/", middlewareAuth, reservationController.CreateReservation)
	reservation.Patch("/staff", middlewareAuth, StaffAuthMiddleware, reservationController.UpdateStaff)
	reservation.Patch("/status", middlewareAuth, StaffAuthMiddleware, reservationController.UpdateStatus)
	reservation.Patch("/payment", middlewareAuth, reservationController.UpdatePayment)

	reservationTask := api.Group("/reservation_task")
	reservationTask.Get("/", middlewareAuth, reservationTaskController.GetAllReservationTask)
	reservationTask.Get("/:reservation_id", middlewareAuth, StaffAuthMiddleware, reservationTaskController.GetReservationTaskByReservationID)
	reservationTask.Post("/", middlewareAuth, StaffAuthMiddleware, reservationTaskController.CreateReservationTask)
	reservationTask.Patch("/staff", middlewareAuth, StaffAuthMiddleware, reservationTaskController.UpdateReservationTaskStaff)
	reservationTask.Patch("/status", middlewareAuth, StaffAuthMiddleware, reservationTaskController.UpdateReservationTaskStatus)

	promotionPrice := api.Group("/promotion_price")
	promotionPrice.Get("/", promotionPriceController.GetAll)
	promotionPrice.Get("/:id", middlewareAuth, StaffAuthMiddleware, promotionPriceController.GetByID)
	promotionPrice.Post("/", middlewareAuth, AdminAuthMiddleware, promotionPriceController.Create)

	room_type_promotion_price := api.Group("/room_type_promotion_price")
	room_type_promotion_price.Get("/room_type/:room_type_id", roomTypePromotionPriceController.GetByRoomTypeID)

	dashboard := api.Group("/dashboard")
	dashboard.Post("/reservation/room_type", dashboardController.GetDashboardRoomTypeReservation)
	dashboard.Post("/reservation/room_type_by_booking", dashboardController.GetRoomTypeReservationCountByBooking)
	dashboard.Post("/reservation/service_type", dashboardController.GetDashboardServiceTypeReservation)
	dashboard.Post("/reservation/service_type_by_booking", dashboardController.GetServiceTypeReservationCountByBooking)
	dashboard.Post("/reservation/payment_type", dashboardController.GetDashboardReservationByPaymentType)
	dashboard.Post("/reservation/maintenance", dashboardController.GetTotalMaintenanceByRoomType)
}
