package project

import (
	"context"
	"time"

	"encore.dev/beta/errs"
	"encore.dev/storage/sqldb"
)

// Project represents a project in the system
type Project struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateProjectParams defines the parameters for creating a new project
type CreateProjectParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Validate validates the create project parameters
func (p *CreateProjectParams) Validate() error {
	if p.Name == "" {
		return &errs.Error{
			Code:    errs.InvalidArgument,
			Message: "project name is required",
		}
	}
	return nil
}

// Database instance for the project service
var db = sqldb.NewDatabase("project", sqldb.DatabaseConfig{
	Migrations: "./migrations",
})

//encore:api public method=POST path=/projects
func CreateProject(ctx context.Context, params *CreateProjectParams) (*Project, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	project := &Project{
		Name:        params.Name,
		Description: params.Description,
		Status:      "ACTIVE",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
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
			Message: "failed to create project",
		}
	}

	return project, nil
}
