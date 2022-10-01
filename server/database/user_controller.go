package database

import (
	"Area/database/models"

	"github.com/davecgh/go-spew/spew"
)

type userController struct {
	User *models.User
}

type UserController interface {
	Create(user *models.User) (*models.User, error)
	Get() (*models.User, error)
	GetById(id uint) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(id uint) (*models.User, error)
}

func (userController) Create(user models.User) (*models.User, error) {
	result := db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	spew.Dump(user)
	return &user, nil
}

func (userController) Get() ([]models.User, error) {
	var users []models.User
	err := db.Model(&models.User{}).Preload("Triggers").Find(&users).Error
	return users, err
}

func (userController) GetById(id uint) (*models.User, error) {
	var user models.User
	err := db.First(&user, id).Error
	return &user, err
}

func (userController) GetByUsername(username string) (*models.User, error) {
	var user models.User
	err := db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (userController) Update(models.User) (*models.User, error) {
	return nil, nil
}

func (userController) Delete(id string) (*models.User, error) {
	return nil, nil
}
