package models


type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"unique;type:varchar(255)"`
	Password string  `json:"password"`
	Products []Product
}
