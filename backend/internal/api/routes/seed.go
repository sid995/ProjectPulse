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

// SeedHandler handles the database seeding route
func SeedHandler(c *gin.Context) {
	// Check if we're in development mode
	if os.Getenv("GIN_MODE") == "release" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Seeding is only available in development mode",
		})
		return
	}

	db := database.GetDB()
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Database connection not available",
		})
		return
	}

	// Begin transaction
	tx := db.Begin()
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to start transaction: %v", tx.Error),
		})
		return
	}

	// Disable foreign key constraints
	if err := tx.Exec("ALTER TABLE users DISABLE TRIGGER ALL").Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to disable user constraints: %v", err),
		})
		return
	}
	if err := tx.Exec("ALTER TABLE teams DISABLE TRIGGER ALL").Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to disable team constraints: %v", err),
		})
		return
	}

	// Delete existing data
	if err := deleteAllData(tx); err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to delete existing data: %v", err),
		})
		return
	}

	// Define the order of seeding to handle dependencies
	seedFiles := []string{
		"team.json",           // First: Teams
		"user.json",           // Second: Users
		"project.json",        // Third: Projects
		"task.json",           // Fourth: Tasks
		"taskAssignment.json", // Fifth: TaskAssignments
		"projectTeam.json",    // Sixth: ProjectTeams
		"comment.json",        // Seventh: Comments
		"attachment.json",     // Eighth: Attachments
	}

	seedDir := "pkg/database/seed_data"

	// Seed all tables
	for _, filename := range seedFiles {
		filePath := filepath.Join(seedDir, filename)
		if err := seedTable(tx, filePath, filename); err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Failed to seed table %s: %v", filename, err),
			})
			return
		}
		log.Printf("Successfully seeded table from %s", filename)
	}

	// Update team managers after all data is created
	if err := updateTeamManagers(tx); err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to update team managers: %v", err),
		})
		return
	}

	// Re-enable foreign key constraints
	if err := tx.Exec("ALTER TABLE users ENABLE TRIGGER ALL").Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to re-enable user constraints: %v", err),
		})
		return
	}
	if err := tx.Exec("ALTER TABLE teams ENABLE TRIGGER ALL").Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to re-enable team constraints: %v", err),
		})
		return
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to commit transaction: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Database seeded successfully",
	})
}

// TeamUpdate represents the structure for updating team managers
type TeamUpdate struct {
	TeamName             string `json:"teamName"`
	ProductOwnerUserID   int    `json:"productOwnerUserId"`
	ProjectManagerUserID int    `json:"projectManagerUserId"`
}

// updateTeamManagers updates teams with product owner and project manager IDs
func updateTeamManagers(tx *gorm.DB) error {
	// Define team updates
	teamUpdates := []TeamUpdate{
		{
			TeamName:             "Quantum Innovations",
			ProductOwnerUserID:   1, // Assuming AliceJones is ID 1
			ProjectManagerUserID: 2, // Assuming BobSmith is ID 2
		},
		{
			TeamName:             "Nebula Research",
			ProductOwnerUserID:   3, // Assuming CarolWhite is ID 3
			ProjectManagerUserID: 4, // Assuming DaveBrown is ID 4
		},
		{
			TeamName:             "Orion Solutions",
			ProductOwnerUserID:   5, // Assuming EveClark is ID 5
			ProjectManagerUserID: 1, // Reusing AliceJones
		},
		{
			TeamName:             "Krypton Developments",
			ProductOwnerUserID:   2, // Reusing BobSmith
			ProjectManagerUserID: 3, // Reusing CarolWhite
		},
		{
			TeamName:             "Zenith Technologies",
			ProductOwnerUserID:   4, // Reusing DaveBrown
			ProjectManagerUserID: 5, // Reusing EveClark
		},
	}

	for _, update := range teamUpdates {
		result := tx.Model(&models.Team{}).
			Where("team_name = ?", update.TeamName).
			Updates(map[string]interface{}{
				"product_owner_user_id":   update.ProductOwnerUserID,
				"project_manager_user_id": update.ProjectManagerUserID,
			})

		if result.Error != nil {
			return fmt.Errorf("failed to update team %s: %v", update.TeamName, result.Error)
		}
	}

	return nil
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

	switch tableName {
	case "user.json":
		var users []models.User
		if err := json.Unmarshal(data, &users); err != nil {
			return fmt.Errorf("failed to parse users: %v", err)
		}
		// Create users with their original team IDs
		for _, user := range users {
			if err := tx.Create(&user).Error; err != nil {
				return fmt.Errorf("failed to create user %s: %v", user.Username, err)
			}
		}

	case "team.json":
		var teams []models.Team
		if err := json.Unmarshal(data, &teams); err != nil {
			return fmt.Errorf("failed to parse teams: %v", err)
		}
		if err := tx.Create(&teams).Error; err != nil {
			return fmt.Errorf("failed to create teams: %v", err)
		}

	case "project.json":
		var projects []models.Project
		if err := json.Unmarshal(data, &projects); err != nil {
			return fmt.Errorf("failed to parse projects: %v", err)
		}
		if err := tx.Create(&projects).Error; err != nil {
			return fmt.Errorf("failed to create projects: %v", err)
		}

	case "task.json":
		var tasks []models.Task
		if err := json.Unmarshal(data, &tasks); err != nil {
			return fmt.Errorf("failed to parse tasks: %v", err)
		}
		if err := tx.Create(&tasks).Error; err != nil {
			return fmt.Errorf("failed to create tasks: %v", err)
		}

	case "comment.json":
		var comments []models.Comment
		if err := json.Unmarshal(data, &comments); err != nil {
			return fmt.Errorf("failed to parse comments: %v", err)
		}
		if err := tx.Create(&comments).Error; err != nil {
			return fmt.Errorf("failed to create comments: %v", err)
		}

	case "attachment.json":
		var attachments []models.Attachment
		if err := json.Unmarshal(data, &attachments); err != nil {
			return fmt.Errorf("failed to parse attachments: %v", err)
		}
		if err := tx.Create(&attachments).Error; err != nil {
			return fmt.Errorf("failed to create attachments: %v", err)
		}

	case "projectTeam.json":
		var projectTeams []models.ProjectTeam
		if err := json.Unmarshal(data, &projectTeams); err != nil {
			return fmt.Errorf("failed to parse projectTeams: %v", err)
		}
		if err := tx.Create(&projectTeams).Error; err != nil {
			return fmt.Errorf("failed to create projectTeams: %v", err)
		}

	case "taskAssignment.json":
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

func createTeamsWithoutUsers(tx *gorm.DB, filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read team seed file: %v", err)
	}

	var teams []models.Team
	if err := json.Unmarshal(data, &teams); err != nil {
		return fmt.Errorf("failed to parse teams: %v", err)
	}

	// Create teams without user associations
	for _, team := range teams {
		// Clear user references
		team.ProductOwnerUserID = nil
		team.ProjectManagerUserID = nil

		if err := tx.Omit("Users").Create(&team).Error; err != nil {
			return fmt.Errorf("failed to create team %s: %v", team.TeamName, err)
		}
	}

	return nil
}
