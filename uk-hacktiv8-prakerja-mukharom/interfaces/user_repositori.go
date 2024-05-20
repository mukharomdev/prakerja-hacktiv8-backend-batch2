package interfaces

import "uk-hacktiv8-prakerja-mukharom/models"

type IUserRepository interface {
	Store(p model.User) (model.User, error)
	Delete(id uint) (model.User, error)
	Update(p model.User) error
	Get(id uint) (model.User, error)
	GetAll() ([]model.User, error)
}
