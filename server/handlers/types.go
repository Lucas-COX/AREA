package handlers

import (
	"Area/database/models"
	"time"
)

// Bodies for /login and /register routes
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
	Token string                 `json:"token"`
}

type ReactionRequestBody struct {
	Type   models.ReactionType       `json:"type"`
	Action models.ReactionActionType `json:"action"`
	Token  string                    `json:"token"`
}

// Body for /trigger route
type TriggerRequestBody struct {
	Title       string              `json:"title"`
	Description string              `json:"description"`
	Action      ActionRequestBody   `json:"action"`
	Reaction    ReactionRequestBody `json:"reaction"`
	Active      bool                `json:"active"`
}

type ActionResponseBody struct {
	ID        uint                   `json:"id"`
	Type      models.ActionType      `json:"type"`
	Event     models.ActionEventType `json:"event"`
	TriggerID uint                   `json:"trigger_id"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
}

type ReactionResponseBody struct {
	ID        uint                   `json:"id"`
	Type      models.ActionType      `json:"type"`
	Action    models.ActionEventType `json:"action"`
	TriggerID uint                   `json:"trigger_id"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
	Token     string                 `json:"token"`
}

// Bodies for user and trigger getters
type TriggerBody struct {
	ID          uint                 `json:"id"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
	Title       string               `json:"title"`
	Description string               `json:"description"`
	UserID      uint                 `json:"user_id"`
	Active      bool                 `json:"active"`
	Action      ActionResponseBody   `json:"action"`
	Reaction    ReactionResponseBody `json:"reaction"`
}

type UserBody struct {
	ID           uint          `json:"id"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	Username     string        `json:"username"`
	Triggers     []TriggerBody `json:"triggers"`
	GoogleLogged bool          `json:"google_logged"`
}
