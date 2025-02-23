package routes

import (
	"backend/internal/models"
	"backend/pkg/database"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reflect"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SeedHandler(c *gin.Context) {
	if os.Getenv("GIN_MODE") == "release" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Seeding is only available in development mode",
		})
	}
	db := database.GetDB()

	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Database connection not available",
		})
		return
	}

	seedDir := "pkg/database/seed_data"

	seedFiles := []string{
		"user.json",
		"team.json",
		"project.json",
		"task.json",
		"comment.json",
		"attachment.json",
		"projectTeam.json",
		"taskAssignment.json",
	}

	tx := db.Begin()
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to begin transaction",
		})
		return
	}

	if err := deleteAllData(tx); err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete data",
		})
		return
	}

	for _, filename := range seedFiles {
		filePath := filepath.Join(seedDir, filename)
		tableName := filename[:len(filename)-5]

		if err := seedTable(tx, filePath, tableName); err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to seed table " + tableName,
			})
			return
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to commit transaction",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Database seeded successfully",
	})
}

func deleteAllData(tx *gorm.DB) error {
	// Delete in correct order to handle foreign key constraints
	models := []interface{}{
		&models.Comment{},        // No dependencies
		&models.Attachment{},     // No dependencies
		&models.TaskAssignment{}, // No dependencies
		&models.ProjectTeam{},    // No dependencies
		&models.Task{},           // Depends on Projects and Users
		&models.Project{},        // Depends on Teams
		&models.Team{},           // Depends on Users
		&models.User{},           // Base table
	}

	for _, model := range models {
		tableName := tx.NamingStrategy.TableName(reflect.TypeOf(model).Elem().Name())
		if err := tx.Exec(fmt.Sprintf("DELETE FROM %s", tableName)).Error; err != nil {
			return fmt.Errorf("failed to delete from %s: %v", tableName, err)
		}
		log.Printf("Deleted all records from %s", tableName)
	}

	log.Println("All data deleted successfully")
	return nil
}

func seedTable(tx *gorm.DB, filePath string, tableName string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read seed file %s: %v", filePath, err)
	}

	// Define the seeding order based on dependencies
	switch tableName {
	case "user.json": // Base table - no dependencies
		var users []models.User
		if err := json.Unmarshal(data, &users); err != nil {
			return fmt.Errorf("failed to parse users: %v", err)
		}
		if err := tx.Create(&users).Error; err != nil {
			return fmt.Errorf("failed to create users: %v", err)
		}

	case "team.json": // Depends on Users
		var teams []models.Team
		if err := json.Unmarshal(data, &teams); err != nil {
			return fmt.Errorf("failed to parse teams: %v", err)
		}
		if err := tx.Create(&teams).Error; err != nil {
			return fmt.Errorf("failed to create teams: %v", err)
		}

	case "project.json": // Independent
		var projects []models.Project
		if err := json.Unmarshal(data, &projects); err != nil {
			return fmt.Errorf("failed to parse projects: %v", err)
		}
		if err := tx.Create(&projects).Error; err != nil {
			return fmt.Errorf("failed to create projects: %v", err)
		}

	case "task.json": // Depends on Projects and Users
		var tasks []models.Task
		if err := json.Unmarshal(data, &tasks); err != nil {
			return fmt.Errorf("failed to parse tasks: %v", err)
		}
		if err := tx.Create(&tasks).Error; err != nil {
			return fmt.Errorf("failed to create tasks: %v", err)
		}

	case "comment.json": // Depends on Tasks and Users
		var comments []models.Comment
		if err := json.Unmarshal(data, &comments); err != nil {
			return fmt.Errorf("failed to parse comments: %v", err)
		}
		if err := tx.Create(&comments).Error; err != nil {
			return fmt.Errorf("failed to create comments: %v", err)
		}

	case "attachment.json": // Depends on Tasks and Users
		var attachments []models.Attachment
		if err := json.Unmarshal(data, &attachments); err != nil {
			return fmt.Errorf("failed to parse attachments: %v", err)
		}
		if err := tx.Create(&attachments).Error; err != nil {
			return fmt.Errorf("failed to create attachments: %v", err)
		}

	case "projectTeam.json": // Depends on Projects and Teams
		var projectTeams []models.ProjectTeam
		if err := json.Unmarshal(data, &projectTeams); err != nil {
			return fmt.Errorf("failed to parse projectTeams: %v", err)
		}
		if err := tx.Create(&projectTeams).Error; err != nil {
			return fmt.Errorf("failed to create projectTeams: %v", err)
		}

	case "taskAssignment.json": // Depends on Tasks and Users
		var taskAssignments []models.TaskAssignment
		if err := json.Unmarshal(data, &taskAssignments); err != nil {
			return fmt.Errorf("failed to parse taskAssignments: %v", err)
		}
		if err := tx.Create(&taskAssignments).Error; err != nil {
			return fmt.Errorf("failed to create taskAssignments: %v", err)
		}

	default:
		return fmt.Errorf("unknown table name: %s", tableName)
	}

	log.Printf("Table %s seeded successfully", tableName)
	return nil
}
