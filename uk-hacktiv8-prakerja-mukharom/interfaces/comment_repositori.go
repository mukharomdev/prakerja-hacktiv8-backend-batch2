package interfaces

import "uk-hacktiv8-prakerja-mukharom/models"

type ICommentRepository interface {
	Store(p model.Comment) (model.Comment, error)
	Delete(id uint) (model.Comment, error)
	Update(p model.Comment) error
	Get(id uint) (model.Comment, error)
	GetAll() ([]model.Comment, error)
}
