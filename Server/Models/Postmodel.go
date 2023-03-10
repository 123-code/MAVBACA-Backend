package Models

import (
    "gorm.io/gorm"
)

type Reserve struct {
    gorm.Model
    ID    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
    Phone string `json:"phone"`
}