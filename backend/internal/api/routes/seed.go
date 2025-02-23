// package routes

// import (
// 	"backend/internal/models"
// 	"backend/pkg/database"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"os"
// 	"path/filepath"

// 	"github.com/gin-gonic/gin"
// )

// // SeedHandler handles the database seeding route
// func SeedHandler(c *gin.Context) {
// 	// Check if we're in development mode
// 	if os.Getenv("GIN_MODE") == "release" {
// 		c.JSON(http.StatusForbidden, gin.H{
// 			"error": "Seeding is only available in development mode",
// 		})
// 		return
// 	}

// 	db := database.GetDB()
// 	if db == nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "Database connection not available",
// 		})
// 		return
// 	}

// 	// Read and parse all seed files
// 	seedDir := "pkg/database/seed_data"

// 	// Define the order of seeding to handle dependencies
// 	seedFiles := []string{
// 		"user.json",
// 		"team.json",
// 		"project.json",
// 		"task.json",
// 		"comment.json",
// 		"attachment.json",
// 		"projectTeam.json",
// 		"taskAssignment.json",
// 	}

// 	// Begin transaction
// 	tx := db.Begin()
// 	if tx.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "Failed to start transaction",
// 		})
// 		return
// 	}

// 	for _, filename := range seedFiles {
// 		filepath := filepath.Join(seedDir, filename)
// 		data, err := ioutil.ReadFile(filepath)
// 		if err != nil {
// 			tx.Rollback()
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"error": fmt.Sprintf("Failed to read seed file %s: %v", filename, err),
// 			})
// 			return
// 		}

// 		switch filename {
// 		case "user.json":
// 			var users []models.User
// 			if err := json.Unmarshal(data, &users); err != nil {
// 				tx.Rollback()
// 				c.JSON(http.StatusInternalServerError, gin.H{
// 					"error": fmt.Sprintf("Failed to parse users: %v", err),
// 				})
// 				return
// 			}
// 			for _, user := range users {
// 				if err := tx.Create(&user).Error; err != nil {
// 					tx.Rollback()
// 					c.JSON(http.StatusInternalServerError, gin.H{
// 						"error": fmt.Sprintf("Failed to create user: %v", err),
// 					})
// 					return
// 				}
// 			}

// 		case "team.json":
// 			var teams []models.Team
// 			if err := json.Unmarshal(data, &teams); err != nil {
// 				tx.Rollback()
// 				c.JSON(http.StatusInternalServerError, gin.H{
// 					"error": fmt.Sprintf("Failed to parse teams: %v", err),
// 				})
// 				return
// 			}
// 			for _, team := range teams {
// 				if err := tx.Create(&team).Error; err != nil {
// 					tx.Rollback()
// 					c.JSON(http.StatusInternalServerError, gin.H{
// 						"error": fmt.Sprintf("Failed to create team: %v", err),
// 					})
// 					return
// 				}
// 			}

// 		case "project.json":
// 			var projects []models.Project
// 			if err := json.Unmarshal(data, &projects); err != nil {
// 				tx.Rollback()
// 				c.JSON(http.StatusInternalServerError, gin.H{
// 					"error": fmt.Sprintf("Failed to parse projects: %v", err),
// 				})
// 				return
// 			}
// 			for _, project := range projects {
// 				if err := tx.Create(&project).Error; err != nil {
// 					tx.Rollback()
// 					c.JSON(http.StatusInternalServerError, gin.H{
// 						"error": fmt.Sprintf("Failed to create project: %v", err),
// 					})
// 					return
// 				}
// 			}

// 		case "task.json":
// 			var tasks []models.Task
// 			if err := json.Unmarshal(data, &tasks); err != nil {
// 				tx.Rollback()
// 				c.JSON(http.StatusInternalServerError, gin.H{
// 					"error": fmt.Sprintf("Failed to parse tasks: %v", err),
// 				})
// 				return
// 			}
// 			for _, task := range tasks {
// 				if err := tx.Create(&task).Error; err != nil {
// 					tx.Rollback()
// 					c.JSON(http.StatusInternalServerError, gin.H{
// 						"error": fmt.Sprintf("Failed to create task: %v", err),
// 					})
// 					return
// 				}
// 			}

// 		case "comment.json":
// 			var comments []models.Comment
// 			if err := json.Unmarshal(data, &comments); err != nil {
// 				tx.Rollback()
// 				c.JSON(http.StatusInternalServerError, gin.H{
// 					"error": fmt.Sprintf("Failed to parse comments: %v", err),
// 				})
// 				return
// 			}
// 			for _, comment := range comments {
// 				if err := tx.Create(&comment).Error; err != nil {
// 					tx.Rollback()
// 					c.JSON(http.StatusInternalServerError, gin.H{
// 						"error": fmt.Sprintf("Failed to create comment: %v", err),
// 					})
// 					return
// 				}
// 			}

// 		case "attachment.json":
// 			var attachments []models.Attachment
// 			if err := json.Unmarshal(data, &attachments); err != nil {
// 				tx.Rollback()
// 				c.JSON(http.StatusInternalServerError, gin.H{
// 					"error": fmt.Sprintf("Failed to parse attachments: %v", err),
// 				})
// 				return
// 			}
// 			for _, attachment := range attachments {
// 				if err := tx.Create(&attachment).Error; err != nil {
// 					tx.Rollback()
// 					c.JSON(http.StatusInternalServerError, gin.H{
// 						"error": fmt.Sprintf("Failed to create attachment: %v", err),
// 					})
// 					return
// 				}
// 			}
// 		}
// 	}

// 	// Commit transaction
// 	if err := tx.Commit().Error; err != nil {
// 		tx.Rollback()
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": fmt.Sprintf("Failed to commit transaction: %v", err),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Database seeded successfully",
// 	})
// }

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
	models := []interface{}{
		&models.User{},
		&models.Team{},
		&models.Project{},
		&models.Task{},
		&models.Comment{},
		&models.Attachment{},
		&models.ProjectTeam{},
		&models.TaskAssignment{},
	}

	for _, model := range models {
		tableName := tx.NamingStrategy.TableName(reflect.TypeOf(model).Elem().Name())
		if err := tx.Exec(fmt.Sprintf("DELETE FROM %s", tableName)).Error; err != nil {
			return err
		}
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
		tx.Create(&users)

	case "team.json":
		var teams []models.Team
		if err := json.Unmarshal(data, &teams); err != nil {
			return fmt.Errorf("failed to parse teams: %v", err)
		}
		tx.Create(&teams)

	case "project.json":
		var projects []models.Project
		if err := json.Unmarshal(data, &projects); err != nil {
			return fmt.Errorf("failed to parse projects: %v", err)
		}
		tx.Create(&projects)

	case "task.json":
		var tasks []models.Task
		if err := json.Unmarshal(data, &tasks); err != nil {
			return fmt.Errorf("failed to parse tasks: %v", err)
		}
		tx.Create(&tasks)

	case "comment.json":
		var comments []models.Comment
		if err := json.Unmarshal(data, &comments); err != nil {
			return fmt.Errorf("failed to parse comments: %v", err)
		}
		tx.Create(&comments)

	case "attachment.json":
		var attachments []models.Attachment
		if err := json.Unmarshal(data, &attachments); err != nil {
			return fmt.Errorf("failed to parse attachments: %v", err)
		}
		tx.Create(&attachments)

	case "projectTeam.json":
		var projectTeams []models.ProjectTeam
		if err := json.Unmarshal(data, &projectTeams); err != nil {
			return fmt.Errorf("failed to parse projectTeams: %v", err)
		}
		tx.Create(&projectTeams)

	case "taskAssignment.json":
		var taskAssignments []models.TaskAssignment
		if err := json.Unmarshal(data, &taskAssignments); err != nil {
			return fmt.Errorf("failed to parse taskAssignments: %v", err)
		}
		tx.Create(&taskAssignments)
	default:
		return fmt.Errorf("unknown table name: %s", tableName)
	}
	log.Println("Table seeded successfully:", tableName)
	return nil
}
