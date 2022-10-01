package handlers

import (
	"Area/database"
	"Area/database/models"
	"context"
	"errors"

	"github.com/go-chi/jwtauth/v5"
)

func UserFromContext(ctx context.Context) (*models.User, error) {
	_, claims, err := jwtauth.FromContext(ctx)

	if err != nil {
		return nil, err
	}
	if id, exists := claims["id"].(float64); exists {
		user, err := database.User.GetById(uint(id), true)
		return user, err
	}
	return nil, errors.New("bad token")
}
