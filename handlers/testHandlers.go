package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/njeru-codes/Go-JWTS/database"
	"github.com/njeru-codes/Go-JWTS/dto"
	model "github.com/njeru-codes/Go-JWTS/models"
	"github.com/njeru-codes/Go-JWTS/utils"
)

func GetProfileHandler(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var user model.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(fiber.Map{
		"id":    user.ID,
		"email": user.Email,
	})
}

func LoginHandler(c *fiber.Ctx) error {
	var req dto.LoginUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if errors := utils.ValidateStruct(&req); errors != nil {
		return c.Status(400).JSON(fiber.Map{
			"errors": errors,
		})
	}

	var user model.User
	if err := database.DB.First(&user, "email = ?", req.Email).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}
	if !utils.CheckPassword(user.Password, req.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
		"user": fiber.Map{
			"id":    user.ID,
			"email": user.Email,
		},
	})
}

func RegisterUserHandler(c *fiber.Ctx) error {
	var req dto.RegisterNewUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}
	// validate body
	if errors := utils.ValidateStruct(&req); errors != nil {
		return c.Status(400).JSON(fiber.Map{
			"errors": errors,
		})
	}

	// hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": "internal server error",
		})
	}

	user := model.User{
		Email:    req.Email,
		Password: string(hashedPassword),
	}
	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": "internal server error",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "user created",
		"user": fiber.Map{
			"id":    user.ID,
			"email": user.Email,
		},
	})
}
