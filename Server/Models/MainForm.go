package Models

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type MainForm struct {
    gorm.Model
    ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    FirstName   string    `json:"firstname"`
	LastName    string    `json:"lastname"`
	Phone       string    `json:"phone"`
    Email       string    `json:"email"`
	Contact     string    `json:"contact"`
	Residence   string    `json:"residence"`
	Adults      int        `json:"adults"`
	Children    int        `json:"children"`
    Comments	string    `json:"comments"`
}

func (p *MainForm) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New()
	return nil
}