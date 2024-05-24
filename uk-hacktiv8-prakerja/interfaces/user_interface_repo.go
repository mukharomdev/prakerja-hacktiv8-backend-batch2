package interfaces

import  "uk-hacktiv8-prakerja/models"

type IUserRepository interface {
	Store(p models.User) (models.User, error)
	Delete(id uint) (models.User, error)
	Update(p models.User) error
	Get(id uint) (models.User, error)
	GetAll() ([]models.User, error)
	IsExist(id uint)(bool)
}