package models

import "gorm.io/gorm"

type Trigger struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string
	UserID      uint `gorm:"not null"`
	User        User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (t *Trigger) TableName() string { return "triggers" }
