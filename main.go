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

	fiber := server.NewFiberServer(usecase, repository)
	go fiber.Start()

	<-signals
	fmt.Println("Server is shutting down")
	if err := fiber.Close(); err != nil {
		log.Fatal("Server is not shutting down", err)
	}
	
	if err := db.Close(); err != nil {
		log.Fatal("MySQL is not shutting down", err)
	}

	fmt.Println("Server was successful shutdown")
}

func initRepository(
	mysql *sqlx.DB,
) *domain.Repository {
	return &domain.Repository{
		UserRepository:    repository.NewUserRepository(mysql),
		SessionRepository: repository.NewSessionRepository(mysql),
	}
}

func initUsecase(
	cfg *config.Config,
	repo *domain.Repository,
) *domain.Usecase {
	googleUsecase := usecase.NewGoogleUsecase(cfg)
	userUsecase := usecase.NewUserUsecase(repo.UserRepository)
	sessionUsecase := usecase.NewSessionUsecase(repo.SessionRepository)
	authUsecase := usecase.NewAuthUsecase(googleUsecase, userUsecase, sessionUsecase)

	return &domain.Usecase{
		AuthUsecase:   authUsecase,
		GoogleUsecase: googleUsecase,
		UserUsecase:   userUsecase,
	}
}
