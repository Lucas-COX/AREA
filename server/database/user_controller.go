package database

import (
	"Area/database/models"
	"Area/lib"
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

func (userController) Create(user models.User) error {
	err := db.Create(&user).Error
	return err
}

func (userController) Get() ([]models.User, error) {
	return nil, nil
}

func (userController) GetById(id uint) (*models.User, error) {
	var user models.User
	err := db.First(&user, id).Error
	lib.CheckError(err)
	return &user, nil
}

func (userController) GetByUsername(username string) (*models.User, error) {
	var user models.User = models.User{Username: username}
	err := db.First(&user).Error
	lib.CheckError(err)
	return &user, nil
}

func (userController) Update(models.User) (*models.User, error) {
	return nil, nil
}

func (userController) Delete(id string) (*models.User, error) {
	return nil, nil
}
