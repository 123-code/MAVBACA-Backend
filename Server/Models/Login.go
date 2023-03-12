package Models

import (
    "gorm.io/gorm"
	"github.com/google/uuid"
)

type Login struct {
    gorm.Model
    ID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    UserName  string `json:"name"`
    Password string `json:"password"`
}

func (p *Login) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New()
	return nil
}