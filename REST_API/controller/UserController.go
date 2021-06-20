package controller

import (
	"fmt"

	"github.com/CCNIT1998/OCG/REST_API/model"
	repo "github.com/CCNIT1998/OCG/REST_API/repository"
	"github.com/gofiber/fiber/v2"
)

func GetAllUser(c *fiber.Ctx) error {
	return c.JSON(repo.UserRepo.GetAllUserRepo())
}

func FindUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		// url params id err
		return c.Status(400).SendString(err.Error()) 
	}

	user, err := repo.UserRepo.FindUserByIdRepo(int64(id))
	if err != nil {
		// not found id
		return c.Status(404).SendString(err.Error())
	}

	// find id in data
	return c.JSON(user)
}


func CreateUser(c *fiber.Ctx) error {
	user := new(model.User)
	err := c.BodyParser(&user)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	userId := repo.UserRepo.CreateNewUser(user)
	return c.SendString(fmt.Sprintf("New user is created successfully with id = %d", userId))
}


func UpsertUser(c *fiber.Ctx) error{
	user := new(model.User)
	err := c.BodyParser(&user)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	id := repo.UserRepo.UpsertUserRepo(user)
	return c.SendString(fmt.Sprintf("User with id = %d is successfully upserted", id))
}


func UpdateUser(c *fiber.Ctx) error{
	updatedUser := new(model.User)

	err := c.BodyParser(&updatedUser)
	// if error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	err = repo.UserRepo.UpdateUserRepo(updatedUser)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	return c.SendString(fmt.Sprintf("User with id = %d is successfully updated", updatedUser.Id))
}


func DeleteUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	err2 := repo.UserRepo.DeleteUserkByIdRepo(int64(id))
	if err2 != nil {
		return c.Status(404).SendString(err2.Error())
	} else {
		return c.SendString(fmt.Sprintf("delete user with id = %d successfully", id))
	}
}