package models

import (
	"database/sql/driver"

	"gorm.io/gorm"
)

type ReactionType string

<<<<<<< HEAD
type ReactionActionType string
=======
type ReactionAction string
>>>>>>> 73b5e7a (feat(server): add action and reaction models)

const (
	Discord ReactionType = "discord"
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
<<<<<<< HEAD
	Type      ReactionType       `gorm:"not null"`
	Action    ReactionActionType `gorm:"not null"`
	TriggerID uint
	Token     string `gorm:"not null"`
=======
	Type      ReactionType   `gorm:"not null"`
	Action    ReactionAction `gorm:"not null"`
	TriggerID uint
>>>>>>> 73b5e7a (feat(server): add action and reaction models)
}

func (r *Reaction) TableName() string { return "reactions" }
