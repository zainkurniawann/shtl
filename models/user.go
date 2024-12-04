// models/user.go
package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Status    string `json:"status"`

	// RoleID (foreign key) untuk relasi dengan Role
	RoleID   uint   `json:"role_id"`
	RoleCode string `json:"role_code"` // Menyimpan kode role (misalnya "SA" untuk Superadmin)

	// Relasi Many-to-Many dengan Role
	Roles []Role `gorm:"many2many:user_roles;"`
}