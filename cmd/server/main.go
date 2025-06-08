package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	log.Println("Setting up server...")

	// create an fiber app
	app := fiber.New(fiber.Config{
		AppName: "FlipFlop API",
	})

	// middlewares
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // allow all origins
		AllowMethods: "GET,POST,PUT,DELETE,PATCH",
		AllowHeaders: "Origin,Content-Type,Accept,Authorization",
	}))

	// health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "UP",
			"message": "Server is running smoothly!",
		})
	})

	port := os.Getenv("PORT")
	if port ==  "" {
		port = "8080" // default port
	}

	log.Printf("Starting server on port %s...\n", port)
	// start the server
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}

}
