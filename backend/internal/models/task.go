package models

import "time"

// Task represents the task table
type Task struct {
	ID             int        `gorm:"primaryKey;autoIncrement" json:"id"`
	Title          string     `json:"title"`
	Description    *string    `json:"description,omitempty"`
	Status         *string    `json:"status,omitempty"`
	Priority       *string    `json:"priority,omitempty"`
	Tags           *string    `json:"tags,omitempty"`
	StartDate      *time.Time `json:"startDate,omitempty"`
	DueDate        *time.Time `json:"dueDate,omitempty"`
	Points         *int       `json:"points,omitempty"`
	ProjectID      int        `gorm:"not null" json:"projectId"`
	AuthorUserID   int        `gorm:"not null" json:"authorUserId"`
	AssignedUserID *int       `json:"assignedUserId,omitempty"`

	Project         Project          `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE" json:"project"`
	Author          User             `gorm:"foreignKey:AuthorUserID;constraint:OnDelete:CASCADE" json:"author"`
	Assignee        *User            `gorm:"foreignKey:AssignedUserID;constraint:OnDelete:SET NULL" json:"assignee"`
	TaskAssignments []TaskAssignment `json:"taskAssignments,omitempty"`
	Attachments     []Attachment     `json:"attachments,omitempty"`
	Comments        []Comment        `json:"comments,omitempty"`
}

// TaskAssignment represents the many-to-many relationship between Users and Tasks
type TaskAssignment struct {
	ID     int `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID int `gorm:"not null" json:"userId"`
	TaskID int `gorm:"not null" json:"taskId"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user"`
	Task Task `gorm:"foreignKey:TaskID;constraint:OnDelete:CASCADE" json:"task"`
}
