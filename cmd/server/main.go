package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/njeru-codes/Go-JWTS/config"
	"github.com/njeru-codes/Go-JWTS/database"
	"github.com/njeru-codes/Go-JWTS/middleware"
	"github.com/njeru-codes/Go-JWTS/routes"
	"github.com/njeru-codes/Go-JWTS/utils"
)

func main() {
	// load config
	cfg := config.LoadEnv()
	utils.InitJWT(cfg.JWT_SECRET)
	middleware.InitJWT(cfg.JWT_SECRET)

	// connect the database
	database.ConnectDB(cfg.DATABASE_DSN)

	// define server
	app := fiber.New(fiber.Config{
		AppName:           "Go JWT SERVER v1.0.0",
		EnablePrintRoutes: true,
	})

	// add midlewares
	app.Use(logger.New())
	app.Get("/api/v1/metics", monitor.New()) //to collect usage metrics

	// add routes here
	allRoutes := app.Group("/api/v1")
	routes.SetupTestRoutes(allRoutes)

	// add error middlewares

	//start serrver
	err := app.Listen(":3000")
	if err != nil {
		log.Fatalf("error while stating server: %v", err)
	}

}
