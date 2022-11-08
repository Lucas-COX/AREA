package handlers

import (
	"Area/database/models"
	"time"
)

type authRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type authResponseBody struct {
	Token string `json:"token"`
}

type logoutResponseBody struct {
	Message string `json:"message"`
}

type triggerRequestBody struct {
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

type triggerBody struct {
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

type userBody struct {
	ID        uint            `json:"id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	Username  string          `json:"username"`
	Role      models.UserRole `json:"role"`
	Triggers  []triggerBody   `json:"triggers"`
	Services  []string        `json:"services"`
}

type aboutResponse struct {
	Client client `json:"client"`
	Server server `json:"server"`
}
