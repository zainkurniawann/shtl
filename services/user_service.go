// services/user_service.go
package services

import (
	"errors"
	"shtlpg/models"
	"github.com/jinzhu/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

// GetRoleByCode retrieves a role by its role_code
func (s *UserService) GetRoleByCode(roleCode string, role *models.Role) error {
	if err := s.db.Where("role_code = ?", roleCode).First(role).Error; err != nil {
		return err
	}
	return nil
}

// services/user_service.go

// GetUserByID retrieves a user by their ID
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	// Preload "Role" untuk mengambil data Role bersama dengan User
	if err := s.db.Preload("Role").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}


// CreateUser creates a new user in the database with a role based on role_code
func (s *UserService) CreateUser(user *models.User) (*models.User, error) {
	// Check if email already exists
	var existingUser models.User
	if err := s.db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return nil, errors.New("email already in use")
	}

	// Fetch role based on role_code
	var role models.Role
	if err := s.db.Where("role_code = ?", user.RoleCode).First(&role).Error; err != nil {
		return nil, errors.New("role not found")
	}

	// Assign role ID to user
	user.RoleID = role.ID

	// Create the user
	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUser updates an existing user and their role based on role_code
// UpdateUser updates an existing user with roles
func (s *UserService) UpdateUser(id uint, user *models.User) (*models.User, error) {
	var existingUser models.User
	if err := s.db.Preload("Roles").First(&existingUser, id).Error; err != nil {
		return nil, err
	}

	// Update fields
	existingUser.FirstName = user.FirstName
	existingUser.LastName = user.LastName
	existingUser.Email = user.Email
	existingUser.Phone = user.Phone
	existingUser.Address = user.Address
	existingUser.Status = user.Status

	// Update roles if provided (using RoleCode)
	if user.RoleCode != "" {
		var role models.Role
		if err := s.db.Where("role_code = ?", user.RoleCode).First(&role).Error; err != nil {
			return nil, err
		}
		// Assign the found role to the user
		existingUser.RoleID = role.ID
		existingUser.Roles = append(existingUser.Roles, role)
	}

	if err := s.db.Save(&existingUser).Error; err != nil {
		return nil, err
	}

	return &existingUser, nil
}


// DeleteUser deletes a user by ID
func (s *UserService) DeleteUser(id uint) error {
	if err := s.db.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
