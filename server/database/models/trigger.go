package models

import (
	"time"

	"gorm.io/gorm"
)

type TriggerData struct {
	Timestamp    time.Time `copier:"-"`
	Author       string    `copier:"-"`
	Title        string    `copier:"-"`
	Description  string    `copier:"-"`
	ActionData   string
	ReactionData string
}

type Trigger struct {
	gorm.Model
	Title           string `gorm:"not null"`
	Description     string
	Active          bool `gorm:"default:false"`
	UserID          uint
	User            User
	ActionService   string
	Action          string
	ReactionService string
	Reaction        string `gorm:"not null"`
	Data            []byte
}

func (t *Trigger) TableName() string { return "triggers" }
