package models

import "time"

// Project represents the project table
type Project struct {
	ID          int        `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	StartDate   *time.Time `json:"startDate,omitempty"`
	EndDate     *time.Time `json:"endDate,omitempty"`

	Tasks        []Task        `json:"tasks,omitempty"`
	ProjectTeams []ProjectTeam `json:"projectTeams,omitempty"`
}

// ProjectTeam represents the many-to-many relationship between Project and Team
type ProjectTeam struct {
	ID        int     `gorm:"primaryKey;autoIncrement" json:"id"`
	TeamID    int     `gorm:"not null" json:"teamId"`
	ProjectID int     `gorm:"not null" json:"projectId"`
	Team      Team    `gorm:"foreignKey:TeamID;constraint:OnDelete:CASCADE" json:"team"`
	Project   Project `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE" json:"project"`
}
