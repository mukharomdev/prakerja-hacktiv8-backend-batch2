package interfaces

import "uk-hacktiv8-prakerja-mukharom/models"

type ISocialmediaRepository interface {
	Store(p model.Socialmedia) (model.Socialmedia, error)
	Delete(id uint) (model.Socialmedia, error)
	Update(p model.Socialmedia) error
	Get(id uint) (model.Socialmedia, error)
	GetAll() ([]model.Socialmedia, error)
}
