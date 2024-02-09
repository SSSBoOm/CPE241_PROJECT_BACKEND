package main

import (
	"log"

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

	fiber := server.NewFiberServer(usecase, repository)

	fiber.Start()
}

func initRepository(
	mysql *sqlx.DB,
) *domain.Repository {
	return &domain.Repository{
		User: repository.NewUserRepository(mysql),
	}
}

func initUsecase(
	cfg *config.Config,
	repository *domain.Repository,
) *domain.Usecase {
	googleUsecase := usecase.NewGoogleUsecase(cfg)
	userUsecase := usecase.NewUserUsecase(repository.User)
	authUsecase := usecase.NewAuthUsecase(googleUsecase, userUsecase)

	return &domain.Usecase{
		AuthUsecase:   authUsecase,
		GoogleUsecase: googleUsecase,
		UserUsecase:   userUsecase,
	}
}
