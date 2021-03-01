package rest

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Route : for router
func Route(app *fiber.App, db *gorm.DB) {
	InitBookRoutes(app.Group("/book"), db)
}
