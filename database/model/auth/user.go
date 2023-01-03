package auth

import (
	"github.com/afolabiolayinka/contact-go/database/model"
	"github.com/google/uuid"
)

// User : model
type User struct {
	model.Model
	Email       string    `gorm:"UNIQUE_INDEX"`
	Password    string    `gorm:"Size:256"`
	CreatedByID uuid.UUID `gorm:"Column:created_by" json:"createdBy"`
	UpdatedByID uuid.UUID `gorm:"Column:updated_by" json:"updatedBy"`
	DeletedByID uuid.UUID `gorm:"Column:deleted_by" json:"deletedBy"`
}

// TableName : sets the name of the table
func (User) TableName() string {
	return "auth_users"
}
