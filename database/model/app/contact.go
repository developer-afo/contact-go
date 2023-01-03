package books

import (
	"github.com/afolabiolayinka/contact-go/database/model"

	"github.com/google/uuid"
)

// Contact : model
type Contact struct {
	model.Model
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	CreatedByID uuid.UUID `gorm:"Column:created_by" json:"createdBy"`
	UpdatedByID uuid.UUID `gorm:"Column:updated_by" json:"updatedBy"`
	DeletedByID uuid.UUID `gorm:"Column:deleted_by" json:"deletedBy"`
}

// TableName : sets the name of the table
func (Contact) TableName() string {
	return "app_contacts"
}
