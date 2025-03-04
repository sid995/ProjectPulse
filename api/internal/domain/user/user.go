package domain

import (
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Email regex pattern for basic validation
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

// FullName returns the user's full name
func (u *User) FullName() string {
	return u.FirstName + " " + u.LastName
}

// SetPassword sets the user's password after hashing it
func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// ComparePassword compares the provided password with the user's hashed password
func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// BeforeCreate will be called before creating a user in the database
func (u *User) BeforeCreate(tx *gorm.DB) error {
	// Validate email format
	if !u.isValidEmail(u.Email) {
		return fmt.Errorf("invalid email format")
	}

	// Validate password length
	if len(u.Password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	// Validate role
	if !u.isValidRole(u.Role) {
		return fmt.Errorf("invalid user role")
	}

	return nil
}

// isValidEmail checks if the email format is valid
func (u *User) isValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

// isValidRole checks if the role is valid
func (u *User) isValidRole(role Role) bool {
	switch role {
	case RoleAdmin, RoleProjectManager, RoleDeveloper, RoleViewer:
		return true
	default:
		return false
	}
}
