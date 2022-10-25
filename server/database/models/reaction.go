package models

import (
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

type Reaction struct {
	gorm.Model
	Type   ReactionType       `gorm:"not null"`
	Action ReactionActionType `gorm:"not null"`
}

func (r *Reaction) TableName() string { return "reactions" }
