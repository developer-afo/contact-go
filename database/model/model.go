package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Model : model
type Model struct {
	UUID uuid.UUID `gorm:"PRIMARY_KEY; Type:uuid" json:"UUID"`
	gorm.Model
}

// Prepare : function
func (model *Model) Prepare() {
	uid, _ := uuid.NewRandom()
	model.UUID = uid
}
