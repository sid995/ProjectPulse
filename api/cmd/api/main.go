package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/sid995/projectpulse/internal/config"
	"github.com/sid995/projectpulse/internal/database"
	"github.com/spf13/viper"
)

func main() {
	// Load configuration
	initConfig()

	// Initialize logger
	logger, err := config.NewLogger()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	sugar := logger.Sugar()
	sugar.Info("Starting ProjectPulse API server")

	// Set Gin mode
	if viper.GetString("GO_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Initialize database
	db, err := database.Initialize(logger)
	if err != nil {
		sugar.Fatalf("Failed to initialize database: %v", err)
	}

	// Run database migrations
	if err := database.RunMigrations(db, logger); err != nil {
		sugar.Fatalf("Failed to run database migrations: %v", err)
	}

	// Initialize router
	router := gin.Default()

	// Health check route
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "OK",
			"message": "ProjectPulse API is running",
		})
	})

	// Start server
	port := viper.GetString("server.port")
	if port == "" {
		port = "8080"
	}

	// Start the server in a goroutine
	go func() {
		sugar.Infof("Starting ProjectPulse API server on :%s", port)
		if err := router.Run(":" + port); err != nil {
			sugar.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Set up graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	sugar.Info("Shutting down server...")

	// Close database connection
	if err := database.Close(); err != nil {
		sugar.Errorf("Error closing database connection: %v", err)
	}

	sugar.Info("Server gracefully stopped")
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../..")
	viper.AddConfigPath("config")
	viper.AddConfigPath(".")

	// Set defaults
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("GO_ENV", "development")

	// Read configuration
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Config file not found, using defaults")
		} else {
			log.Fatalf("Error reading config file: %v", err)
		}
	}

	// Read from environment variables
	viper.AutomaticEnv()
}
