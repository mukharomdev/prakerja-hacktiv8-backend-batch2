package repository

import (
	"fmt"

	"uk-hacktiv8-prakerja-mukharom/database"
	"uk-hacktiv8-prakerja-mukharom/models"
)

type CommentRepository struct {
}

func (s *CommentRepository) Store(p models.Comment) (models.Comment, error) {
	db := database.GetDatabase()
	err := db.Create(&p).Error

	if err != nil {
		return models.Comment{}, err
	}

	return p, nil
}

func (s *CommentRepository) Delete(id uint) (models.Comment, error) {
	db := database.GetDatabase()
	n, err := s.Get(id)
	if err != nil {
		return models.Comment{}, err
	}

	err = db.Delete(&n).Error
	if err != nil {
		return models.Comment{}, fmt.Errorf("cannot delete file: %v", err)
	}

	return n, nil
}

func (s *CommentRepository) Update(p models.Comment) error {
	db := database.GetDatabase()

	err := db.Save(&p).Error
	if err != nil {
		return fmt.Errorf("cannot update client on pg: %v", err)
	}

	return nil
}

func (s *CommentRepository) Get(id uint) (models.Comment, error) {
	db := database.GetDatabase()
	var p models.Comment
	err := db.First(&p, id).Error

	if err != nil {
		return models.Comment{}, fmt.Errorf("cannot find Comment by id: %v", err)
	}

	return p, nil
}

func (s *CommentRepository) GetAll() ([]models.Comment, error) {
	db := database.GetDatabase()
	var p []models.Comment
	err := db.Find(&p).Error

	if err != nil {
		return nil, fmt.Errorf("cannot find Comments: %v", err)
	}

	return p, err
}
