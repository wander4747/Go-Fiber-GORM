package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wander4747/gorm/controller"
)

func ProductRoute(route fiber.Router) {
	route.Get("", controller.GetProducts)
	route.Post("", controller.CreateProduc)
	route.Get("/:id", controller.GetProduct)
	route.Put("/:id", controller.UpdateProduct)
	route.Delete("/:id", controller.DeleteProduct)
}
