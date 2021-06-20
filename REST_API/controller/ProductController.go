package controller

import (
	"fmt"

	"github.com/CCNIT1998/OCG/REST_API/model"
	repo "github.com/CCNIT1998/OCG/REST_API/repository"
	"github.com/gofiber/fiber/v2"
)

func CreateNewProduct(c *fiber.Ctx) error {
	product := new(model.Product)
	err := c.BodyParser(&product)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	productId := repo.ProductRepo.CreateNewProductRepo(product)
	return c.SendString(fmt.Sprintf("New user is created successfully with id = %d", productId))
}


func UpsertProduct(c *fiber.Ctx) error{
	product := new(model.Product)
	err := c.BodyParser(&product)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	id := repo.ProductRepo.UpsertProductRepo(product)
	return c.SendString(fmt.Sprintf("Product with id = %d is successfully upserted", id))
}


func UpdateProduct(c *fiber.Ctx) error{
	updatedProduct := new(model.Product)

	// data sent here
	err := c.BodyParser(&updatedProduct)

	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	err = repo.ProductRepo.UpdateProductRepo(updatedProduct)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	return c.SendString(fmt.Sprintf("Product with id = %d is successfully updated", updatedProduct.Id))
}



func DeleteProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	err2 := repo.ProductRepo.DeleteProductkByIdRepo(int64(id))
	if err2 != nil {
		return c.Status(404).SendString(err2.Error())
	} else {
		return c.SendString(fmt.Sprintf("delete product with id = %d successfully", id))
	}
}


func GetAllProduct(c *fiber.Ctx) error {
	return c.JSON(repo.ProductRepo.GetAllProductRepo())
}


func UpdateImageInProduct(c *fiber.Ctx) error{
	id, e := c.ParamsInt("id")
	if e != nil {
		return c.Status(400).SendString(e.Error())
	}

	updatedProduct := new(model.Product)
	err := c.BodyParser(&updatedProduct)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	err = repo.ProductRepo.UpdateProductRepo1(updatedProduct, int64(id))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	return c.SendString(fmt.Sprintf("Product with id = %d is successfully updated", int64(id)))
}
