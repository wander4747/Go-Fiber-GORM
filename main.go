package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/wander4747/gorm/database"
	"github.com/wander4747/gorm/models"
	"github.com/wander4747/gorm/routes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	var err error

	database.Connection, err = gorm.Open(sqlite.Open("product.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection opened to database")
	database.Connection.AutoMigrate(&models.Product{})
}

func main() {

	app := fiber.New()
	app.Use(logger.New())

	api := app.Group("/api")

	routes.ProductRoute(api.Group("/products"))

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Sorry can't find that!")
	})

	log.Fatal(app.Listen(":7779"))
}
