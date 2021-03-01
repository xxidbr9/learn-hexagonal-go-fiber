package rest

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Book :
type Book struct {
	ID   string `json:"id"`
	Name string `json:"book_name"`
}

// Res :
type Res struct {
	Msg  string `json:"msg"`
	Data []Book `json:"data"`
}

// InitBookRoutes :
func InitBookRoutes(app fiber.Router, db *gorm.DB) {

	// var res Res

	data := []Book{}

	data = append(data, Book{ID: "1", Name: "Ini Buku"})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(&Res{Msg: "Bisa coii", Data: data})
	})
}
