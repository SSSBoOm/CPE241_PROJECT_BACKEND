package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/SSSBoOm/CPE241_Project_Backend/db"
	"github.com/SSSBoOm/CPE241_Project_Backend/domain"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/config"
	"github.com/SSSBoOm/CPE241_Project_Backend/repository"
	"github.com/SSSBoOm/CPE241_Project_Backend/server"
	"github.com/SSSBoOm/CPE241_Project_Backend/usecase"
	"github.com/jmoiron/sqlx"
)

//	@title			CPE241 Project Backend API
//	@version		1.0
//	@description	This is a sample server celler server.

// @schemes	https http
// @host		localhost:8080
// @BasePath	/
// @securityDefinitions.apikey 	ApiKeyAuth
// @in 													header
// @name												Authorization
func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Cant Load Config : ", err)
	}

	db, err := db.NewMySQLConnect(cfg.MYSQL_URI)
	if err != nil {
		log.Fatal("Cant Connect To Mysql : ", err)
	}

	repository := initRepository(db)
	usecase := initUsecase(cfg, repository)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	fiber := server.NewFiberServer(cfg, usecase, repository)
	go fiber.Start()

	<-signals
	fmt.Println("Server is shutting down")
	if err := fiber.Close(); err != nil {
		log.Fatal("Server is not shutting down", err)
	}
	fmt.Println("fiber was successful")

	if err := db.Close(); err != nil {
		log.Fatal("MySQL is not shutting down", err)
	}
	fmt.Println("db was successful")
	fmt.Println("Server was successful shutdown")
}

func initRepository(
	mysql *sqlx.DB,
) *domain.Repository {
	return &domain.Repository{
		UserRepository:                   repository.NewUserRepository(mysql),
		SessionRepository:                repository.NewSessionRepository(mysql),
		RoleRepository:                   repository.NewRoleRepository(mysql),
		RoomRepository:                   repository.NewRoomRepository(mysql),
		RoomTypeRepository:               repository.NewRoomTypeRepository(mysql),
		PaymentRepository:                repository.NewPaymentRepository(mysql),
		PaymentTypeRepository:            repository.NewPaymentTypeRepository(mysql),
		MaintenanceRepository:            repository.NewMaintenanceRepository(mysql),
		MaintenanceLogRepository:         repository.NewMaintenanceLogRepository(mysql),
		ReservationRepository:            repository.NewReservationRepository(mysql),
		ReservationTaskRepository:        repository.NewReservationTaskRepository(mysql),
		ServiceRepository:                repository.NewServiceRepository(mysql),
		ServiceTypeRepository:            repository.NewServiceTypeRepository(mysql),
		PromotionPriceRepository:         repository.NewPromotionPriceRepository(mysql),
		RoomTypePromotionPriceRepository: repository.NewRoomTypePromotionPriceRepository(mysql),
	}
}

func initUsecase(
	cfg *config.Config,
	repo *domain.Repository,
) *domain.Usecase {
	googleUsecase := usecase.NewGoogleUsecase(cfg)
	userUsecase := usecase.NewUserUsecase(repo.UserRepository, repo.RoleRepository)
	sessionUsecase := usecase.NewSessionUsecase(repo.SessionRepository)
	authUsecase := usecase.NewAuthUsecase(googleUsecase, userUsecase, sessionUsecase)
	roleUsecase := usecase.NewRoleUsecase(repo.RoleRepository)
	paymentTypeUsecase := usecase.NewPaymentTypeUsecase(repo.PaymentTypeRepository)
	paymentUsecase := usecase.NewPaymentUsecase(repo.PaymentRepository, paymentTypeUsecase)
	roomTypeUsecase := usecase.NewRoomTypeUsecase(repo.RoomTypeRepository, repo.RoomRepository)
	roomUsecase := usecase.NewRoomUsecase(repo.RoomRepository, roomTypeUsecase)
	maintenanceLogUsecase := usecase.NewMaintenanceLogUsecase(repo.MaintenanceLogRepository, userUsecase)
	maintenanceUsecase := usecase.NewMaintenanceUsecase(repo.MaintenanceRepository, maintenanceLogUsecase, roomUsecase)
	serviceUsecase := usecase.NewServiceUsecase(repo.ServiceRepository, repo.ServiceTypeRepository)
	serviceTypeRepository := usecase.NewServiceTypeUsecase(repo.ServiceTypeRepository, serviceUsecase)
	reservationUsecase := usecase.NewReservationUsecase(repo.ReservationRepository, roomTypeUsecase, roomUsecase, serviceUsecase, paymentUsecase, userUsecase)
	reservationTaskUsecase := usecase.NewReservationTaskUseCase(repo.ReservationTaskRepository, reservationUsecase, userUsecase)
	roomTypePromotionPriceUsecase := usecase.NewRoomTypePromotionPriceUsecase(repo.RoomTypePromotionPriceRepository, repo.PromotionPriceRepository, roomTypeUsecase)
	promotionPriceUsecase := usecase.NewPromotionPriceUsecase(repo.PromotionPriceRepository, roomTypePromotionPriceUsecase)

	return &domain.Usecase{
		AuthUsecase:                   authUsecase,
		GoogleUsecase:                 googleUsecase,
		UserUsecase:                   userUsecase,
		SessionUsecase:                sessionUsecase,
		RoleUsecase:                   roleUsecase,
		RoomUsecase:                   roomUsecase,
		RoomTypeUsecase:               roomTypeUsecase,
		PaymentUsecase:                paymentUsecase,
		PaymentTypeUsecase:            paymentTypeUsecase,
		MaintenanceUsecase:            maintenanceUsecase,
		MaintenanceLogUsecase:         maintenanceLogUsecase,
		ReservationUsecase:            reservationUsecase,
		ReservationTaskUsecase:        reservationTaskUsecase,
		ServiceUsecase:                serviceUsecase,
		ServiceTypeUsecase:            serviceTypeRepository,
		RoomTypePromotionPriceUsecase: roomTypePromotionPriceUsecase,
		PromotionPriceUsecase:         promotionPriceUsecase,
	}
}
