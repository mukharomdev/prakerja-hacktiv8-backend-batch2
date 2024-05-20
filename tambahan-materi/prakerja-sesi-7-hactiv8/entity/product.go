package entity

type Product struct {
	ID     uint   `json:"id"`
	Name   string `json:"name" gorm:"type:varchar(255)"`
	Price  int    `json:"price"`
	UserID uint
}
