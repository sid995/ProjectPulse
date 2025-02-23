package models

// Comment represents the comment table
type Comment struct {
	ID     int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Text   string `json:"text"`
	TaskID int    `gorm:"not null" json:"taskId"`
	UserID int    `gorm:"not null" json:"userId"`

	Task Task `gorm:"foreignKey:TaskID;constraint:OnDelete:CASCADE" json:"task"`
	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user"`
}
