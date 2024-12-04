// routes/user_routes.go
package routes

import (
	"shtlpg/controllers"
	"shtlpg/databases"
	"shtlpg/services"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	db := databases.GetDB()
	userService := services.NewUserService(db)
	userController := controllers.NewUserController(userService)

	// Define user routes
	app.Post("/users", userController.CreateUser)
	app.Get("/users/:id", userController.GetUser)
	app.Put("/users/:id", userController.UpdateUser)
	app.Delete("/users/:id", userController.DeleteUser)
}
