package handlers

import (
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

// Body for /trigger route
type TriggerRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Bodies for user and trigger getters
type TriggerBody struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserID      uint      `json:"user_id"`
}

type UserBody struct {
	ID        uint          `json:"id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	Username  string        `json:"username"`
	Triggers  []TriggerBody `json:"triggers"`
}
