package models

import (
	"database/sql/driver"

	"gorm.io/gorm"
)

type ActionType string

type ActionEventType string

const (
	GmailAction   ActionType = "gmail"
	DiscordAction ActionType = "discord"
)

func (a *ActionType) Scan(value interface{}) error {
	*a = ActionType(value.([]byte))
	return nil
}

func (a ActionType) Value() (driver.Value, error) {
	return string(a), nil
}

type Action struct {
	gorm.Model
	Type      ActionType `gorm:"not null"`
	EventType string     `gorm:"not null"`
	TriggerID uint
}

func (a *Action) TableName() string { return "actions" }
