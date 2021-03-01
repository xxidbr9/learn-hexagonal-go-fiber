package main

import (
	"fmt"
	"log"
	"nando/gorm_fiber/config"
	"nando/gorm_fiber/http/rest"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	con := fmt.Sprintf("host=%v user=%v password=%v DB.name=%v port=%v", config.DBHost, config.DBUser, config.DBPass, config.DBName, config.DBPort)
	gormDB, err := gorm.Open(postgres.Open(con), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(map[string]string{
			"msg": "succes on",
		})
	})

	rest.Route(app, gormDB)
	log.Fatal(app.Listen(":3000"))
}
