package rest

import (
	"nando/gorm_fiber/book/bookinfra/bookquery"
	"nando/gorm_fiber/shared"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetData :
func GetData(db *gorm.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		data := bookquery.GetBooks(db)
		return c.Status(200).JSON(shared.NewRes("Success get All Books", data))
	}
}

// GetBookByID :
func GetBookByID(db *gorm.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		data, err := bookquery.GetBookByID(db, c.Params("id"))
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(shared.ErrorRes(err))
		}
		return c.Status(200).JSON(shared.NewRes("Success get Book Detail", data))
	}
}	
