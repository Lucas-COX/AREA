package database

import "Area/database/models"

type triggerController struct{}

type TriggerController interface {
	Create(trigger models.Trigger) (*models.Trigger, error)
	Get() ([]models.Trigger, error)
	GetById(id string) (*models.Trigger, error)
	GetByTitle(title string) (*models.Trigger, error)
	Update(trigger models.Trigger) (*models.Trigger, error)
	Delete(id string) error
}

func (triggerController) Create(trigger models.Trigger) (*models.Trigger, error) {
	err := db.Create(&trigger).Error
	if err != nil {
		return nil, err
	}
	return &trigger, nil
}

func (triggerController) Get() ([]models.Trigger, error) {
	var triggers []models.Trigger
	var err error
	err = db.Model(&models.Trigger{}).Find(&triggers).Error

	return triggers, err
}

func (triggerController) GetById(id uint) (*models.Trigger, error) {
	var trigger models.Trigger
	var err error
	err = db.First(&trigger, id).Error
	return &trigger, err
}

func (triggerController) GetByTitle(title string) (*models.Trigger, error) {
	var trigger models.Trigger
	var err error
	err = db.Where("title = ?", title).First(&trigger).Error
	return &trigger, err
}

func (triggerController) Update(trigger models.Trigger) (*models.Trigger, error) {
	var err error
	err = db.Save(&trigger).Error
	return &trigger, err
}

func (triggerController) Delete(id string) error {
	var err error
	err = db.Delete(&id).Error
	return err
}
