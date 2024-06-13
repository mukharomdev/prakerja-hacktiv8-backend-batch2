package interfaces

import  "uk-hacktiv8-prakerja/models"

type IPhotoRepository interface {
	Store(p models.Photo) (models.Photo, error)
	Delete(id uint) (models.Photo, error)
	Update(p models.Photo) error
	Get(id uint) (models.Photo, error)
	GetAll() ([]models.Photo, error)
	IsExist(id uint)(bool)
}