// models/role.go
package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name     string `json:"name"`
	RoleCode string `json:"role_code" gorm:"unique"` // Unik agar bisa dibedakan
	Users    []User `gorm:"many2many:user_roles;"`
}
