package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/njeru-codes/Go-JWTS/handlers"
	"github.com/njeru-codes/Go-JWTS/middleware"
)

func SetupTestRoutes(router fiber.Router) {

	testRoutes := router.Group("/auth")

	testRoutes.Post("/login", handlers.LoginHandler)
	testRoutes.Post("/register", handlers.RegisterUserHandler)

	testRoutes.Get("/profile", middleware.JWTMiddleware(), handlers.GetProfileHandler)
}
