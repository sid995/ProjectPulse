package models

// Attachment represents the attachment table
type Attachment struct {
	ID           int     `gorm:"primaryKey;autoIncrement" json:"id"`
	FileURL      string  `json:"fileURL"`
	FileName     *string `json:"fileName,omitempty"`
	TaskID       int     `gorm:"not null" json:"taskId"`
	UploadedByID int     `gorm:"not null" json:"uploadedById"`

	Task       Task `gorm:"foreignKey:TaskID;constraint:OnDelete:CASCADE" json:"task"`
	UploadedBy User `gorm:"foreignKey:UploadedByID;constraint:OnDelete:CASCADE" json:"uploadedBy"`
}
