package models

import (
	"time"

	"gorm.io/gorm"
)

type TriggerData struct {
	Timestamp     time.Time
	Author        string
	Title         string
	Description   string
	ActionToken   string
	ReactionToken string
}

type Trigger struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string
	Active      bool `gorm:"default:false"`
	UserID      uint
	User        User
	ActionID    *uint
	Action      Action `gorm:"not null"`
	ReactionID  *uint
	Reaction    Reaction `gorm:"not null"`
	Data        []byte
}

func (t *Trigger) TableName() string { return "triggers" }
