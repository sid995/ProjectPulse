package domain

import (
	"time"

	"gorm.io/gorm"
)

// Role represents a user role in the system
type Role string

const (
	// RoleAdmin represents an administrator with full access
	RoleAdmin Role = "admin"
	// RoleProjectManager represents a project manager
	RoleProjectManager Role = "project_manager"
	// RoleDeveloper represents a developer
	RoleDeveloper Role = "developer"
	// RoleViewer represents a user with view-only access
	RoleViewer Role = "viewer"
)

// User represents a user in the system
type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Email     string         `json:"email" gorm:"uniqueIndex;not null"`
	Password  string         `json:"-" gorm:"not null"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Role      Role           `json:"role" gorm:"type:varchar(20);default:'viewer'"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
