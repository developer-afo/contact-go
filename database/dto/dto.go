package dto

import (
	"time"

	"github.com/google/uuid"
)

/* Afolabi Olayinka */
/* I prefer to use camelCase for my json naming convention because it saves extra bits than using snake-case
 */

//DTO : DTO
type DTO struct {
	ID        uint      `json:"id"`
	UUID      uuid.UUID `json:"uuid"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	//DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
