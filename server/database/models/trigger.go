package models

import "gorm.io/gorm"

type Trigger struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string
	UserID      uint
	Action      Action   `gorm:"not null"`
	Reaction    Reaction `gorm:"not null"`
	Active      bool     `gorm:"default:false"`
}

func (t *Trigger) TableName() string { return "triggers" }
