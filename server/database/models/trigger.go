package models

import "gorm.io/gorm"

type Trigger struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string
	UserID      uint
}

func (t *Trigger) TableName() string { return "triggers" }
