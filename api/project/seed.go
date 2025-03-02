package project

import (
	"context"
	"fmt"
	"time"

	"encore.dev/beta/errs"
)

// SeedProject represents test data for a project
type SeedProject struct {
	Name        string
	Description string
	Status      string
}

var defaultTestProjects = []SeedProject{
	{
		Name:        "Project Alpha",
		Description: "A test project for development",
		Status:      "ACTIVE",
	},
	{
		Name:        "Project Beta",
		Description: "Another test project for QA",
		Status:      "PLANNING",
	},
	{
		Name:        "Project Gamma",
		Description: "A completed test project",
		Status:      "COMPLETED",
	},
}

// SeedProjectParams defines parameters for seeding projects
type SeedProjectParams struct {
	Count    int           // Number of projects to seed (default: len(defaultTestProjects))
	Projects []SeedProject // Optional custom projects to seed
}

//encore:api public method=POST path=/projects/seed
func SeedProjects(ctx context.Context, params *SeedProjectParams) ([]Project, error) {
	// Use default count if not specified
	count := params.Count
	if count == 0 {
		count = len(defaultTestProjects)
	}

	// Use provided projects or defaults
	seedProjects := params.Projects
	if len(seedProjects) == 0 {
		seedProjects = defaultTestProjects
	}

	// Ensure we have enough projects to seed
	if count > len(seedProjects) {
		// Repeat projects if we need more
		originalLen := len(seedProjects)
		for i := 0; i < count-originalLen; i++ {
			project := seedProjects[i%originalLen]
			project.Name = fmt.Sprintf("%s %d", project.Name, i+1)
			seedProjects = append(seedProjects, project)
		}
	}

	// Seed the projects
	projects := make([]Project, 0, count)
	now := time.Now()

	for i := 0; i < count; i++ {
		seed := seedProjects[i]
		project := &Project{
			Name:        seed.Name,
			Description: seed.Description,
			Status:      seed.Status,
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		err := db.QueryRow(ctx, `
			INSERT INTO projects (name, description, status, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING id
		`, project.Name, project.Description, project.Status, project.CreatedAt, project.UpdatedAt).
			Scan(&project.ID)

		if err != nil {
			return nil, &errs.Error{
				Code:    errs.Internal,
				Message: fmt.Sprintf("failed to seed project %s", seed.Name),
			}
		}

		projects = append(projects, *project)
	}

	return projects, nil
}

//encore:api public method=DELETE path=/projects/seed
func ClearTestData(ctx context.Context) error {
	_, err := db.Exec(ctx, `DELETE FROM projects`)
	if err != nil {
		return &errs.Error{
			Code:    errs.Internal,
			Message: "failed to clear test data",
		}
	}
	return nil
}
