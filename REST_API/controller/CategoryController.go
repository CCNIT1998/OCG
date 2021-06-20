package controller

import (
	"fmt"

	"github.com/CCNIT1998/OCG/REST_API/model"
	repo "github.com/CCNIT1998/OCG/REST_API/repository"
	"github.com/gofiber/fiber/v2"
)

func CreateNewCategory(c *fiber.Ctx) error {
	category := new(model.Category)
	err := c.BodyParser(&category)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	categoryId := repo.CategoryRepo.CreateNewCategoryRepo(category)
	return c.SendString(fmt.Sprintf("New categoryId is created successfully with id = %d", categoryId))
}


func UpsertCategory(c *fiber.Ctx) error{
	category := new(model.Category)
	err := c.BodyParser(&category)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	id := repo.CategoryRepo.UpsertCategoryRepo(category)
	return c.SendString(fmt.Sprintf("category with id = %d is successfully upserted", id))
}


func UpdateCategory(c *fiber.Ctx) error{
	updatedcategory := new(model.Category)

	err := c.BodyParser(&updatedcategory)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	err = repo.CategoryRepo.UpdateCategoryRepo(updatedcategory)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	return c.SendString(fmt.Sprintf("Category with id = %d is successfully updated", updatedcategory.Id))
}



func DeleteCategoryById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	err2 := repo.CategoryRepo.DeleteCategoryByIdRepo(int64(id))
	if err2 != nil {
		return c.Status(404).SendString(err2.Error())
	} else {
		return c.SendString(fmt.Sprintf("delete category with id = %d successfully", id))
	}
}


func GetAllCategory(c *fiber.Ctx) error {
	return c.JSON(repo.CategoryRepo.GetAllCategoryRepo())
}
