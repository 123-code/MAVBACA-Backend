package Models

import (
    "gorm.io/gorm"
    "github.com/google/uuid"
)

type User struct {
    gorm.Model
    ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    UserName  string `json:"name"`
    Email string `json:"email"`
    Password string `json:"password"`
}

func (p *User) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New()
	return nil
}