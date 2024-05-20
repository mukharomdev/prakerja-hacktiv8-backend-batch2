package repository

import (
	"fmt"

	"uk-hacktiv8-prakerja-mukharom/database"
	"uk-hacktiv8-prakerja-mukharom/models"
)

type PhotoRepository struct {
}

func (s *PhotoRepository) Store(p models.Photo) (models.Photo, error) {
	db := database.GetDatabase()
	err := db.Create(&p).Error

	if err != nil {
		return models.Photo{}, err
	}

	return p, nil
}

func (s *PhotoRepository) Delete(id uint) (models.Photo, error) {
	db := database.GetDatabase()
	n, err := s.Get(id)
	if err != nil {
		return models.Photo{}, err
	}

	err = db.Delete(&n).Error
	if err != nil {
		return models.Photo{}, fmt.Errorf("cannot delete file: %v", err)
	}

	return n, nil
}

func (s *PhotoRepository) Update(p models.Photo) error {
	db := database.GetDatabase()

	err := db.Save(&p).Error
	if err != nil {
		return fmt.Errorf("cannot update client on pg: %v", err)
	}

	return nil
}

func (s *PhotoRepository) Get(id uint) (models.Photo, error) {
	db := database.GetDatabase()
	var p models.Photo
	err := db.First(&p, id).Error

	if err != nil {
		return models.Photo{}, fmt.Errorf("cannot find Photo by id: %v", err)
	}

	return p, nil
}

func (s *PhotoRepository) GetAll() ([]models.Photo, error) {
	db := database.GetDatabase()
	var p []models.Photo
	err := db.Find(&p).Error

	if err != nil {
		return nil, fmt.Errorf("cannot find Photos: %v", err)
	}

	return p, err
}
