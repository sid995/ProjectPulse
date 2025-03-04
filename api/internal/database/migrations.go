package database

import (
	"fmt"
	"time"

	userDomain "github.com/sid995/projectpulse/internal/domain/user"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// MigrationRecord is used to keep track of applied migrations
type MigrationRecord struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(255);not null;uniqueIndex"`
	AppliedAt time.Time `gorm:"not null"`
}

// RunMigrations executes all migrations in the proper order
func RunMigrations(db *gorm.DB, logger *zap.Logger) error {
	sugar := logger.Sugar()
	sugar.Info("Running database migrations")

	// Ensure migration table exists
	if err := db.AutoMigrate(&MigrationRecord{}); err != nil {
		return fmt.Errorf("failed to create migrations table: %w", err)
	}

	// Register all migrations here in order
	migrations := []struct {
		Name string
		Fn   func(tx *gorm.DB) error
	}{
		{
			Name: "create_initial_schema",
			Fn:   createInitialSchema,
		},
		// Add more migrations here as needed
	}

	// Execute migrations in transaction
	for _, migration := range migrations {
		var record MigrationRecord
		// Check if migration has already been applied
		if result := db.Where("name = ?", migration.Name).First(&record); result.Error == nil {
			sugar.Infof("Migration %s already applied", migration.Name)
			continue
		}

		sugar.Infof("Applying migration: %s", migration.Name)

		// Run migration in a transaction
		if err := db.Transaction(func(tx *gorm.DB) error {
			if err := migration.Fn(tx); err != nil {
				sugar.Errorf("Migration %s failed: %v", migration.Name, err)
				return err
			}

			// Record successful migration
			return tx.Create(&MigrationRecord{
				Name:      migration.Name,
				AppliedAt: time.Now(),
			}).Error
		}); err != nil {
			return fmt.Errorf("migration %s failed: %w", migration.Name, err)
		}

		sugar.Infof("Migration %s completed successfully", migration.Name)
	}

	sugar.Info("All migrations completed successfully")
	return nil
}

// createInitialSchema creates the initial database schema
func createInitialSchema(tx *gorm.DB) error {
	// Create tables
	if err := tx.AutoMigrate(
		&userDomain.User{},
		// Add more domain models here
	); err != nil {
		return err
	}

	return nil
}
