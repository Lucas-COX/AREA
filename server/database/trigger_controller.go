package database

import "Area/database/models"

type triggerController struct{}

type TriggerController interface {
	Create(trigger *models.Trigger) (*models.Trigger, error)
	Get() ([]models.Trigger, error)
	GetById(id string) (*models.Trigger, error)
	GetByTitle(title string) (*models.Trigger, error)
	Update(trigger *models.Trigger) (*models.Trigger, error)
	Delete(id string) (*models.Trigger, error)
}
