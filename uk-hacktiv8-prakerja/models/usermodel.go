package models

import (
	"time"

	//"github.com/google/uuid"
)



type User struct{
	ID 			uint 		`gorm:"primary_key" json:"id,omitempty"`
	Username 	string 		`gorm:"uniqueIndex;not null" json:"username,omitempty"`
	Email 		string 		`gorm:"uniqueIndex;not null;size > 256" json:"email,omitempty"`
	Password 	string 		`gorm:"uniqueIndex;not null" json:"password,omitempty"`
	Age 		uint 		`gorm:"not null;check:age > 8" json:"age,omitempty"`
	Created_at 	time.Time 	`gorm:"autoCreateTime" json:"created_at,omitempty"`
	Updated_at 	time.Time 	`gorm:"autoUpdateTime:mili " json:"updated_at,omitempty"`
}

// REQUEST

// request
type UserLoginReq struct{
	Email 		string 	`json:"email"`
	Username 	string 	`json:"username"`
	Password 	string 		`json:"password"`
	Age 		uint 		`json:"age"`
}



// RESPON

// respon POST /users/login
type UserLoginRes struct{
	Token 	string `json:"token"`
}


// respon PUT /users

type UserLoginUpdateRes struct{
	ID 			uint 		`json:"id"`
	Username 	string 		`json:"username"`
	Email 		string 		`json:"email"`
	Age 		string 		`json:"age"`
	Updated_at 	time.Time 	`json:"updated_at"`

}


// respons DELETE DELETE /users
// request
//● headers: Authorization (Bearer token string)

// respon
type UserLoginDelRes struct{
	Message string `json:"message"`

}


type UserRegisterReq struct{
	Username 	string 		`json:"username,omitempty"`
	Email 		string 		`json:"email,omitempty"`
	Password 	string 		`json:"password,omitempty"`
	Age 		uint 		`json:"age,omitempty"`
}

type UserRegisterRes struct{
	Age 		uint 		`json:"age,omitempty"`
    Email 		string 		`json:"email,omitempty"`
    ID 			uint 		`json:"id,omitempty"`
	Username 	string 		`json:"username,omitempty"`
}


func NewUser(uslog *UserLoginReq) *User {
	return &User{
		Email:uslog.Email,
		Username:  uslog.Username,
		Password:uslog.Password,
		Age:uslog.Age,
	}
}

// func NewUserWithID(id uint, name string, price float32) *User {
// 	return &User{
// 		ID:    id,
// 		Name:  name,
// 		Price: price,
// 	}
// }
