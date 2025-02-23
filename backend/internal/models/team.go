package models

// Team represents the team table
type Team struct {
	ID                   int    `gorm:"primaryKey;autoIncrement" json:"id"`
	TeamName             string `json:"teamName"`
	ProductOwnerUserID   *int   `json:"productOwnerUserId,omitempty"`
	ProjectManagerUserID *int   `json:"projectManagerUserId,omitempty"`

	ProjectTeams []ProjectTeam `json:"projectTeams,omitempty"`
	Users        []User        `json:"users,omitempty"`
}
