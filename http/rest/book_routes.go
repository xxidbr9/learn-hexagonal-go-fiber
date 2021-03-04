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

// InitBookRoutes :
func InitBookRoutes(app fiber.Router, db *gorm.DB) {

	// var res Res

	// data := []Book{}

	// for i := 1; i <= 10; {
	// 	data = append(data, Book{ID: fmt.Sprint(i), Name: fmt.Sprintf("Ini Buku ke : %d", i)})
	// 	i++
	// }

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.Status(fiber.StatusOK).JSON(shared.NewRes("Success Get All Books", data))
	// })

	// Handling CRUD IS BELOW HERE
	app.Get("/", GetData(db))
	app.Get("/:id", GetBookByID(db))
}
