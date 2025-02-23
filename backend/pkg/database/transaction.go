package database

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

// TxFn is a function that will be called with a database transaction
type TxFn func(tx *gorm.DB) error

// WithTransaction executes a function within a database transaction
func WithTransaction(ctx context.Context, db *gorm.DB, fn TxFn) error {
	tx := db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return fmt.Errorf("failed to begin transaction: %v", tx.Error)
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r) // re-throw panic after rollback
		}
	}()

	if err := fn(tx); err != nil {
		if rbErr := tx.Rollback().Error; rbErr != nil {
			return fmt.Errorf("error: %v, rollback error: %v", err, rbErr)
		}
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

// Example usage:
/*
func CreateUserWithTeam(ctx context.Context, user *models.User, team *models.Team) error {
    return WithTransaction(ctx, GetDB(), func(tx *gorm.DB) error {
        // Create team
        if err := tx.Create(team).Error; err != nil {
            return err
        }

        // Set team ID for user
        user.TeamID = &team.ID

        // Create user
        if err := tx.Create(user).Error; err != nil {
            return err
        }

        return nil
    })
}
*/
