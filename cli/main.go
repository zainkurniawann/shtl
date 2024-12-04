// cli/main.go
package main

import (
	"log"
	"shtlpg/databases"
	"shtlpg/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize the PostgreSQL connection
	if err := databases.ConnectPostgres(); err != nil {
		log.Fatalf("Could not connect to PostgreSQL: %v", err)
	}
	defer databases.CloseDB()

	// Create a new Fiber app
	app := fiber.New()

	// Set up user routes
	routes.UserRoutes(app)

	// Start the server
	log.Println("Server started on http://localhost:8080")
	if err := app.Listen(":1234"); err != nil {
		log.Fatal(err)
	}
}
