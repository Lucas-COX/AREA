package database

import (
	"Area/database/models"
)

type triggerController struct{}

type TriggerController interface {
	Create(trigger *models.Trigger) (*models.Trigger, error)
	Get(user_id uint, all bool) ([]models.Trigger, error)
	GetById(id uint, user_id uint, all bool) (*models.Trigger, error)
	GetByTitle(title string, user_id uint, all bool) (*models.Trigger, error)
	GetActive(all bool) ([]models.Trigger, error)
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

func (triggerController) Get(user_id uint, all bool) ([]models.Trigger, error) {
	var triggers []models.Trigger
	tx := db.Model(&models.Trigger{}).Where("user_id = ?", user_id)
	if all {
		tx = tx.Preload("Action").Preload("Reaction")
	}
	err := tx.Find(&triggers).Error

	return triggers, err
}

func (triggerController) GetById(id uint, user_id uint, all bool) (*models.Trigger, error) {
	var trigger models.Trigger
	tx := db.Model(models.Trigger{}).Where("user_id = ?", user_id)
	if all {
		tx = tx.Preload("Action").Preload("Reaction")
	}
	err := tx.First(&trigger, id).Error
	return &trigger, err
}

func (triggerController) GetByTitle(title string, user_id uint, all bool) (*models.Trigger, error) {
	var trigger models.Trigger

	tx := db.Model(models.Trigger{}).Where("user_id = ?, title = ?", user_id, title)
	if all {
		tx = tx.Preload("Action").Preload("Reaction")
	}
	err := tx.First(&trigger).Error
	return &trigger, err
}

func (triggerController) GetActive(all bool) ([]models.Trigger, error) {
	var triggers []models.Trigger

	tx := db.Model(models.Trigger{}).Where("active = ?", 1).Preload("User")
	if all {
		tx = tx.Preload("Action").Preload("Reaction")
	}
	err := tx.Find(&triggers).Error
	return triggers, err
}

func (triggerController) Update(trigger *models.Trigger) (*models.Trigger, error) {
	err := db.Save(trigger).Error
	if err != nil {
		return nil, err
	}
	err = db.Save(&trigger.Action).Error
	if err != nil {
		return nil, err
	}
	err = db.Save(&trigger.Reaction).Error
	return trigger, err
}

func (triggerController) Delete(trigger *models.Trigger) error {
	err := db.Delete(trigger).Error
	return err
}
