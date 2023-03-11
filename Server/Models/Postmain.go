package Models

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Postmain struct {
    gorm.Model
    ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    Name        string    `json:"name"`
    Email       string    `json:"email"`
    Age         int       `json:"age"`
    Nationality string    `json:"nationality"`
}

func (p *Postmain) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New()
	return nil
}