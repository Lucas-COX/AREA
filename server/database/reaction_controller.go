package database

import "Area/database/models"

type reactionController struct{}

type ReactionController interface {
	Create(reaction *models.Reaction) (*models.Reaction, error)
	Get() ([]models.Reaction, error)
	GetById(id uint) (*models.Reaction, error)
	Update(reaction *models.Reaction) (*models.Reaction, error)
	Delete(reaction *models.Reaction) error
}

func (reactionController) Create(reaction *models.Reaction) (*models.Reaction, error) {
	err := db.Create(reaction).Error
	if err != nil {
		return nil, err
	}
	return reaction, nil
}

func (reactionController) Get() ([]models.Reaction, error) {
	var reactions []models.Reaction
	err := db.Model(models.Reaction{}).Find(&reactions).Error

	return reactions, err
}

func (reactionController) GetById(id uint) (*models.Reaction, error) {
	var reaction models.Reaction
	err := db.Model(&models.Reaction{}).Where("id = ?", id).First(&reaction).Error

	return &reaction, err
}

func (reactionController) Update(reaction *models.Reaction) (*models.Reaction, error) {
	err := db.Save(reaction).Error

	return reaction, err
}

func (reactionController) Delete(reaction *models.Reaction) error {
	err := db.Delete(reaction).Error

	return err
}
