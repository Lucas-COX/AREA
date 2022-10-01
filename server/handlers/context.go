package handlers

import (
	"Area/database"
	"Area/database/models"
	"context"
	"errors"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-chi/jwtauth/v5"
)

func UserFromContext(ctx context.Context) (*models.User, error) {
	_, claims, err := jwtauth.FromContext(ctx)

	if err != nil {
		return nil, err
	}
	spew.Dump(claims)
	if id, exists := claims["id"].(float64); exists {
		user, err := database.User.GetById(uint(id))
		return user, err
	}
	return nil, errors.New("bad token")
}
