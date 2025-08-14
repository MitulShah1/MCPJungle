package model

import "gorm.io/gorm"

// UserRole represents the role of a user in the MCPJungle system.
type UserRole string

const UserRoleAdmin UserRole = "admin"

// User represents a user in the MCPJungle system
type User struct {
	gorm.Model

	Username    string   `gorm:"unique; not null" json:"username"`
	Role        UserRole `gorm:"not null"         json:"role"`
	AccessToken string   `gorm:"unique; not null" json:"access_token"`
}
