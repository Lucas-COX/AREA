package database

import "Area/database/models"

type actionController struct{}

type ActionController interface {
	Create(action *models.Action) (*models.Action, error)
	Get() ([]models.Action, error)
	GetById(id uint) (*models.Reaction, error)
	Update(action *models.Action) (*models.Action, error)
	Delete(action *models.Action) error
}

func (actionController) Create(action *models.Action) (*models.Action, error) {
	err := db.Create(action).Error
	if err != nil {
		return nil, err
	}
	return action, nil
}

func (actionController) Get() ([]models.Action, error) {
	var actions []models.Action
	err := db.Model(models.Action{}).Find(&actions).Error

	return actions, err
}

func (actionController) GetById(id uint) (*models.Reaction, error) {
	var action models.Action
	err := db.Model(models.Action{}).Where("id = ?", id).Find(&action).Error

	return nil, err
}

func (actionController) Update(action *models.Action) (*models.Action, error) {
	err := db.Save(action).Error

	return action, err
}

func (actionController) Delete(action *models.Action) error {
	err := db.Delete(action).Error

	return err
}
