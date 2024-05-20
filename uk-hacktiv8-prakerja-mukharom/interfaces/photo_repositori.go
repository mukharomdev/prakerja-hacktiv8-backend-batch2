package interfaces

import "uk-hacktiv8-prakerja-mukharom/models"

type IPhotoRepository interface {
	Store(p model.Photo) (model.Photo, error)
	Delete(id uint) (model.Photo, error)
	Update(p model.Photo) error
	Get(id uint) (model.Photo, error)
	GetAll() ([]model.Photo, error)
}
