package database

import (
	"Area/database/models"
)

type triggerController struct{}

type TriggerController interface {
	Create(trigger models.Trigger) (*models.Trigger, error)
	Get(user_id uint) ([]models.Trigger, error)
	GetById(id uint, user_id uint) (*models.Trigger, error)
	GetByTitle(title string, user_id uint) (*models.Trigger, error)
	GetActive() ([]models.Trigger, error)
	Update(trigger *models.Trigger) (*models.Trigger, error)
	Delete(trigger *models.Trigger) error
}

func (triggerController) Create(trigger models.Trigger) (*models.Trigger, error) {
	err := db.Create(&trigger).Error
	if err != nil {
		return nil, err
	}
	return &trigger, nil
}

func (triggerController) Get(user_id uint) ([]models.Trigger, error) {
	var triggers []models.Trigger
	err := db.Model(&models.Trigger{}).Where("user_id = ?", user_id).Find(&triggers).Error

	return triggers, err
}

func (triggerController) GetById(id uint, user_id uint) (*models.Trigger, error) {
	var trigger models.Trigger

	err := db.Model(models.Trigger{}).Where("user_id = ?", user_id).First(&trigger, id).Error
	return &trigger, err
}

func (triggerController) GetByTitle(title string, user_id uint) (*models.Trigger, error) {
	var trigger models.Trigger

	err := db.Model(models.Trigger{}).Where("user_id = ?, title = ?", user_id, title).First(&trigger).Error
	return &trigger, err
}

func (triggerController) GetActive() ([]models.Trigger, error) {
	var triggers []models.Trigger

	err := db.Model(models.Trigger{}).Where("active = ?", true).Preload("User").Find(&triggers).Error
	return triggers, err
}

func (triggerController) Update(trigger *models.Trigger) (*models.Trigger, error) {
	err := db.Save(trigger).Error
	if err != nil {
		return nil, err
	}
	return trigger, nil
}

func (triggerController) Delete(trigger *models.Trigger) error {
	err := db.Delete(trigger).Error
	return err
}
