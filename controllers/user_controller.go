// controllers/user_controller.go
package controllers

import (
	"net/http"
	"shtlpg/models"
	"shtlpg/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

// CreateUser handles the creation of a new user
// controllers/user_controller.go
func (ctrl *UserController) CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Fetch the role based on role_code
	var role models.Role
	if err := ctrl.userService.GetRoleByCode(user.RoleCode, &role); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid role code"})
	}

	// Assign the role ID to the user
	user.RoleID = role.ID

	// Now create the user
	createdUser, err := ctrl.userService.CreateUser(&user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(createdUser)
}



// GetUser retrieves a user by ID with role
func (ctrl *UserController) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	userIDUint := uint(userID)

	// Get user with role
	user, err := ctrl.userService.GetUserByID(userIDUint)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.Status(http.StatusOK).JSON(user)
}

// UpdateUser updates a user's details with role
// controllers/user_controller.go
func (ctrl *UserController) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	userIDUint := uint(userID)

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Fetch the role based on role_code for update
	var role models.Role
	if err := ctrl.userService.GetRoleByCode(user.RoleCode, &role); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid role code"})
	}

	// Assign the new role ID to the user
	user.RoleID = role.ID

	updatedUser, err := ctrl.userService.UpdateUser(userIDUint, &user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(updatedUser)
}



// DeleteUser deletes a user by ID
func (ctrl *UserController) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	userIDUint := uint(userID)

	err = ctrl.userService.DeleteUser(userIDUint)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Could not delete user"})
	}

	return c.SendStatus(http.StatusNoContent)
}
