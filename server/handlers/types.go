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

type TriggerRequestBody struct {
	Title           string `json:"title"`
	Description     string `json:"description"`
	ActionService   string `json:"action_service"`
	Action          string `json:"action"`
	ReactionService string `json:"reaction_service"`
	Reaction        string `json:"reaction"`
	ActionData      string `json:"action_data"`
	ReactionData    string `json:"reaction_data"`
	Active          *bool  `json:"active"`
}

type ActionResponseBody struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ReactionResponseBody struct {
	ID          uint      `json:"id"`
	Name        string    `json:"type"`
	Description string    `json:"action"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TriggerBody struct {
	ID              uint      `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	UserID          uint      `json:"user_id"`
	Active          bool      `json:"active"`
	ActionService   string    `json:"action_service"`
	Action          string    `json:"action"`
	ReactionService string    `json:"reaction_service"`
	Reaction        string    `json:"reaction"`
	ActionData      string    `json:"action_data"`
	ReactionData    string    `json:"reaction_data"`
}

type UserBody struct {
	ID        uint            `json:"id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	Username  string          `json:"username"`
	Role      models.UserRole `json:"role"`
	Triggers  []TriggerBody   `json:"triggers"`
	Services  []string        `json:"services"`
}
