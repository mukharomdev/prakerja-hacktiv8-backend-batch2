package models

import(
	"gorm.io/gorm"
	"errors"
)

type Product struct {
	ID uint 	`json:"id" gorm:"primaryKey;autoIncrement,unique"`
	Name string	`json:"name" gorm:"not null"`
	Price uint 	`json:"price" gorm:"not null"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {

	if p.Price == 0 {
		err = errors.New("price is required")
		return
	}

	return
}