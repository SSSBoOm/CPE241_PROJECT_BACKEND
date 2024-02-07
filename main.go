package main

import (
	"log"

	"github.com/SSSBoOm/CPE241_Project_Backend/db"
	"github.com/SSSBoOm/CPE241_Project_Backend/internal/config"
	"github.com/SSSBoOm/CPE241_Project_Backend/route"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	env := config.LoadEnv()

	db, err := db.NewMySQLConnect(env.MYSQL_URI)
	if err != nil {
		log.Fatal("Cant Connect To Mysql : ", err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"massage": "API sBay Data",
		})
	})

	route.MainRoute(app, db, env)

	app.Listen(":8080")
}
