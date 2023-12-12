package controller

import (
	"template-starter-mysql/app/repository"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserRepo *repository.UserRepository
}

func NewUserController(userRepo *repository.UserRepository) *UserController{
	return &UserController{UserRepo: userRepo}
}

func(service *UserController) GetAllUsers(c *fiber.Ctx) error{
	users, err := service.UserRepo.GetAllUsers()
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal Server Error"})
	}

	return c.JSON(users)
}