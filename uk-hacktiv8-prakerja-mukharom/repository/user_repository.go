package repository

import (
	"fmt"

	"uk-hacktiv8-prakerja-mukharom/database"
	"uk-hacktiv8-prakerja-mukharom/models"
)

type UserRepository struct {
}

func (s *UserRepository) Store(p models.User) (models.User, error) {
	db := database.GetDatabase()
	err := db.Create(&p).Error

	if err != nil {
		return models.User{}, err
	}

	return p, nil
}

func (s *UserRepository) Delete(id uint) (models.User, error) {
	db := database.GetDatabase()
	n, err := s.Get(id)
	if err != nil {
		return models.User{}, err
	}

	err = db.Delete(&n).Error
	if err != nil {
		return models.User{}, fmt.Errorf("cannot delete file: %v", err)
	}

	return n, nil
}

func (s *UserRepository) Update(p models.User) error {
	db := database.GetDatabase()

	err := db.Save(&p).Error
	if err != nil {
		return fmt.Errorf("cannot update client on pg: %v", err)
	}

	return nil
}

func (s *UserRepository) Get(id uint) (models.User, error) {
	db := database.GetDatabase()
	var p models.User
	err := db.First(&p, id).Error

	if err != nil {
		return models.User{}, fmt.Errorf("cannot find User by id: %v", err)
	}

	return p, nil
}

func (s *UserRepository) GetAll() ([]models.User, error) {
	db := database.GetDatabase()
	var p []models.User
	err := db.Find(&p).Error

	if err != nil {
		return nil, fmt.Errorf("cannot find Users: %v", err)
	}

	return p, err
}
