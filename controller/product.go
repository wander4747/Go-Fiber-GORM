package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/wander4747/gorm/database"
	"github.com/wander4747/gorm/models"
)

func GetProducts(c *fiber.Ctx) error {
	var products []models.Product
	database.Connection.Find(&products)
	return c.JSON(products)
}

func GetProduct(c *fiber.Ctx) error {
	var product models.Product

	ID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": "id needs to be an integer"})
	}

	result := database.Connection.First(&product, ID)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{})
	}
	return c.JSON(product)
}

func CreateProduc(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err})
	}

	if err := product.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	result := database.Connection.Create(&product)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error})
	}

	return c.JSON(&product)
}

func UpdateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err})
	}

	ID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": "id needs to be an integer"})
	}

	if err := product.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var p models.Product
	result := database.Connection.First(&p, ID)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{})
	}

	database.Connection.Model(p).Updates(product)

	return c.Status(fiber.StatusOK).JSON(product)
}

func DeleteProduct(c *fiber.Ctx) error {

	ID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": "id needs to be an integer"})
	}

	var product models.Product
	result := database.Connection.First(&product, ID)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{})
	}
	database.Connection.Delete(&product)
	return c.Status(fiber.StatusNoContent).JSON(product)
}
