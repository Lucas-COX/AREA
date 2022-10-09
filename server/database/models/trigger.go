package models

import (
	"time"

	"gorm.io/gorm"
)

type TriggerData struct {
	Timestamp   time.Time
	Author      string
	Title       string
	Description string
}

type Trigger struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string
	Active      bool `gorm:"default:false"`
	UserID      uint
	User        User
	Action      Action   `gorm:"not null"`
	Reaction    Reaction `gorm:"not null"`
	Data        []byte
}

func (t *Trigger) TableName() string { return "triggers" }
