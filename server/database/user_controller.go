package database

import (
	"Area/database/models"
	"context"
	"errors"

	"github.com/go-chi/jwtauth/v5"
)

type userController struct {
}

type UserController interface {
	Create(user models.User) (*models.User, error)
	Get(loadTriggers bool) (*models.User, error)
	GetById(id uint, loadTriggers bool) (*models.User, error)
	GetByUsername(username string, loadTriggers bool) (*models.User, error)
	Update(user models.User) (*models.User, error)
	Delete(id uint) error
}

func (userController) Create(user models.User) (*models.User, error) {
	err := db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (userController) Get(loadTriggers bool) ([]models.User, error) {
	var users []models.User
	var err error
	if loadTriggers {
		err = db.Model(&models.User{}).Preload("Triggers").Find(&users).Error
	} else {
		err = db.Model(&models.User{}).Find(&users).Error
	}
	return users, err
}

func (userController) GetFromContext(ctx context.Context) (*models.User, error) {
	_, claims, err := jwtauth.FromContext(ctx)

	if err != nil {
		return nil, err
	}
	if id, exists := claims["id"].(float64); exists {
		user, err := User.GetById(uint(id), true)
		return user, err
	}
	return nil, errors.New("bad token")
}

func (userController) GetById(id uint, loadTriggers bool) (*models.User, error) {
	var user models.User
	var err error
	if loadTriggers {
		err = db.Preload("Triggers").First(&user, id).Error
	} else {
		err = db.First(&user, id).Error
	}
	return &user, err
}

func (userController) GetByUsername(username string, loadTriggers bool) (*models.User, error) {
	var user models.User
	var err error
	if loadTriggers {
		err = db.Where("username = ?", username).Preload("Triggers").First(&user).Error
	} else {
		err = db.Where("username = ?", username).First(&user).Error
	}
	return &user, err
}

func (userController) Update(user models.User) (*models.User, error) {
	err := db.Save(&user).Error
	return &user, err
}

func (userController) Delete(id string) error {
	err := db.Delete(&id).Error
	return err
}
