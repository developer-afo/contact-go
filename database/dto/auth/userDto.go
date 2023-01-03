package auth

import (
	"github.com/afolabiolayinka/contact-go/database/dto"
	"github.com/google/uuid"
)

// UserDTO : model
type UserDTO struct {
	dto.DTO
	Email       string    `json:"email"`
	Password    string    `json:"-"`
	CreatedByID uuid.UUID `json:"createdBy"`
	UpdatedByID uuid.UUID `json:"updatedBy"`
	DeletedByID uuid.UUID `json:"deletedBy"`
}
