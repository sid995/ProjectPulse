package database

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

// Initialize sets up the database connection
func Initialize(logger *zap.Logger) (*gorm.DB, error) {
	dbLogger := logger.Sugar()

	// Build DSN from config
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.name"),
		viper.GetString("database.sslmode"),
	)

	// Open database connection
	dbLogger.Infof("Connecting to database: %s:%s/%s",
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.name"),
	)

	// Configure GORM
	gormConfig := &gorm.Config{}

	// Connect to database
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	sqlDB.SetMaxIdleConns(viper.GetInt("database.max_idle_conns"))
	sqlDB.SetMaxOpenConns(viper.GetInt("database.max_open_conns"))
	sqlDB.SetConnMaxLifetime(viper.GetDuration("database.conn_max_lifetime"))

	// Set global DB instance
	DB = db

	dbLogger.Info("Database connection established")

	return db, nil
}

// Close closes the database connection
func Close() error {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}

// Migrate runs database migrations
func Migrate(models ...interface{}) error {
	if DB == nil {
		return fmt.Errorf("database not initialized")
	}

	return DB.AutoMigrate(models...)
}
