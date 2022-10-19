package handlers

import (
	"Area/database/models"
	"time"
)

type AuthRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponseBody struct {
	Token string `json:"token"`
}

type LogoutResponseBody struct {
	Message string `json:"message"`
}

type ActionRequestBody struct {
	Type  models.ActionType      `json:"type"`
	Event models.ActionEventType `json:"event"`
}

type ReactionRequestBody struct {
	Type   models.ReactionType       `json:"type"`
	Action models.ReactionActionType `json:"action"`
}

type TriggerRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ActionID    *uint  `json:"action_id"`
	ReactionID  *uint  `json:"reaction_id"`
	Active      bool   `json:"active"`
}

type ActionResponseBody struct {
	ID        uint                   `json:"id"`
	Type      models.ActionType      `json:"type"`
	Event     models.ActionEventType `json:"event"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
}

type ReactionResponseBody struct {
	ID        uint                   `json:"id"`
	Type      models.ActionType      `json:"type"`
	Action    models.ActionEventType `json:"action"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
}

type TriggerBody struct {
	ID           uint                 `json:"id"`
	CreatedAt    time.Time            `json:"created_at"`
	UpdatedAt    time.Time            `json:"updated_at"`
	Title        string               `json:"title"`
	Description  string               `json:"description"`
	UserID       uint                 `json:"user_id"`
	Active       bool                 `json:"active"`
	ActionID     *uint                `json:"action_id"`
	Action       ActionResponseBody   `json:"action,omitempty"`
	ReactionID   *uint                `json:"reaction_id"`
	Reaction     ReactionResponseBody `json:"reaction,omitempty"`
	ActionData   string               `json:"action_data"`
	ReactionData string               `json:"reaction_data"`
}

type TriggerSmallBody struct {
	ID           uint                 `json:"id"`
	CreatedAt    time.Time            `json:"created_at"`
	UpdatedAt    time.Time            `json:"updated_at"`
	Title        string               `json:"title"`
	Description  string               `json:"description"`
	UserID       uint                 `json:"user_id"`
	Active       bool                 `json:"active"`
	ActionID     uint                 `json:"action_id"`
	Action       ActionResponseBody   `json:"-"`
	ReactionID   uint                 `json:"reaction_id"`
	Reaction     ReactionResponseBody `json:"-"`
	ActionData   string               `json:"action_data"`
	ReactionData string               `json:"reaction_data"`
}

type UserBody struct {
	ID           uint          `json:"id"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	Username     string        `json:"username"`
	Triggers     []TriggerBody `json:"triggers"`
	GoogleLogged bool          `json:"google_logged"`
}
