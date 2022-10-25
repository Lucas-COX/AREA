package models

import (
	"database/sql/driver"

	"gorm.io/gorm"
)

type ActionType string

type ActionEventType string

const (
	NoneAction    ActionType = "none"
	GmailAction   ActionType = "gmail"
	DiscordAction ActionType = "discord"
)

const (
	NoneEvent    ActionEventType = "none"
	SendEvent    ActionEventType = "send"
	ReceiveEvent ActionEventType = "receive"
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
	Type  ActionType      `gorm:"not null"`
	Event ActionEventType `gorm:"not null"`
}

func (a *Action) TableName() string { return "actions" }
