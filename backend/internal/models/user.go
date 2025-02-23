package models

// User represents the user table
type User struct {
	UserID            int     `gorm:"primaryKey;autoIncrement" json:"userId"`
	CognitoID         string  `gorm:"unique;not null" json:"cognitoId"`
	Username          string  `gorm:"unique;not null" json:"username"`
	ProfilePictureURL *string `json:"profilePictureUrl,omitempty"`
	TeamID            *int    `json:"teamId,omitempty"`

	AuthoredTasks   []Task           `gorm:"foreignKey:AuthorUserID;constraint:OnDelete:CASCADE" json:"authoredTasks,omitempty"`
	AssignedTasks   []Task           `gorm:"foreignKey:AssignedUserID;constraint:OnDelete:SET NULL" json:"assignedTasks,omitempty"`
	TaskAssignments []TaskAssignment `json:"taskAssignments,omitempty"`
	Attachments     []Attachment     `gorm:"foreignKey:UploadedByID;constraint:OnDelete:CASCADE" json:"attachments,omitempty"`
	Comments        []Comment        `json:"comments,omitempty"`
	Team            *Team            `gorm:"foreignKey:TeamID;constraint:OnDelete:SET NULL" json:"team,omitempty"`
}
