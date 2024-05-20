package entity

type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email" gorm:"unique;type:varchar(255)"`
	Password string
	Products []Product
}
