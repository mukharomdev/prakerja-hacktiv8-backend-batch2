package repository

import (
	"fmt"

	"uk-hacktiv8-prakerja-mukharom/database"
	"uk-hacktiv8-prakerja-mukharom/models"
)

type SocialMediaRepository struct {
}

func (s *SocialMediaRepository) Store(p models.SocialMedia) (models.SocialMedia, error) {
	db := database.GetDatabase()
	err := db.Create(&p).Error

	if err != nil {
		return models.SocialMedia{}, err
	}

	return p, nil
}

func (s *SocialMediaRepository) Delete(id uint) (models.SocialMedia, error) {
	db := database.GetDatabase()
	n, err := s.Get(id)
	if err != nil {
		return models.SocialMedia{}, err
	}

	err = db.Delete(&n).Error
	if err != nil {
		return models.SocialMedia{}, fmt.Errorf("cannot delete file: %v", err)
	}

	return n, nil
}

func (s *SocialMediaRepository) Update(p models.SocialMedia) error {
	db := database.GetDatabase()

	err := db.Save(&p).Error
	if err != nil {
		return fmt.Errorf("cannot update client on pg: %v", err)
	}

	return nil
}

func (s *SocialMediaRepository) Get(id uint) (models.SocialMedia, error) {
	db := database.GetDatabase()
	var p models.SocialMedia
	err := db.First(&p, id).Error

	if err != nil {
		return models.SocialMedia{}, fmt.Errorf("cannot find SocialMedia by id: %v", err)
	}

	return p, nil
}

func (s *SocialMediaRepository) GetAll() ([]models.SocialMedia, error) {
	db := database.GetDatabase()
	var p []models.SocialMedia
	err := db.Find(&p).Error

	if err != nil {
		return nil, fmt.Errorf("cannot find SocialMedias: %v", err)
	}

	return p, err
}
