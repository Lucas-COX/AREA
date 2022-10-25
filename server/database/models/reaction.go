package models

import (
	"database/sql/driver"

	"gorm.io/gorm"
)

type ReactionType string

type ReactionActionType string

const (
	NoneReaction    ReactionType = "none"
	DiscordReaction ReactionType = "discord"
)

const (
	NoneReactionAction ReactionActionType = "none"
	SendReaction       ReactionActionType = "send"
)

func (e *ReactionType) Scan(value interface{}) error {
	*e = ReactionType(value.([]byte))
	return nil
}

func (e ReactionType) Value() (driver.Value, error) {
	return string(e), nil
}

type Reaction struct {
	gorm.Model
	Type   ReactionType       `gorm:"not null"`
	Action ReactionActionType `gorm:"not null"`
}

func (r *Reaction) TableName() string { return "reactions" }
