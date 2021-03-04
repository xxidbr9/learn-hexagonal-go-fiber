package main

import (
	"fmt"
	"log"
	"nando/gorm_fiber/book"
	"nando/gorm_fiber/config"
	"nando/gorm_fiber/http/rest"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	con := fmt.Sprintf("host=%v user=%v password=%v database=%v port=%v", config.DBHost, config.DBUser, config.DBPass, config.DBName, config.DBPort)
	gormDB, err := gorm.Open(postgres.Open(con), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migrating all table

	// Rest START
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(map[string]string{
			"msg": "succes on",
		})
	})
	if bookService := book.NewBookApp(gormDB); bookService != nil {
		log.Fatal("Book Broken")
	}
	rest.Route(app, gormDB)
	log.Fatal(app.Listen(":5000"))
}
