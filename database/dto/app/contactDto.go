package books

import (
	"github.com/afolabiolayinka/contact-go/database/dto"

	"github.com/google/uuid"
)

// ContactDTO : DTO
type ContactDTO struct {
	dto.DTO
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	CreatedByID uuid.UUID `json:"createdBy"`
	UpdatedByID uuid.UUID `json:"updatedBy"`
	DeletedByID uuid.UUID `json:"deletedBy"`
}
