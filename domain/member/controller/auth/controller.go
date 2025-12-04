package controller

import (
	"apirest/domain/user/model"
	"apirest/database"
	"apirest/utils"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var req authRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	user := model.User{
		Email:        req.Email,
		Password: utils.GeneratePassword(req.Password),
	}
	res := database.DB.Create(&user)
	if res.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": res.Error.Error(),
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"message": "user created",
	})
}

func Login(c *fiber.Ctx) error {
	var req authRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	var user model.User
	res := database.DB.Where("email = ?", req.Email).First(&user)
	if res.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "user not found",
		})
	}
	if !utils.ComparePassword(user.Password, req.Password) {
		return c.Status(400).JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	token, err := utils.GenerateToken(uint(user.ID))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"token": token,
	})
}

type authRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
