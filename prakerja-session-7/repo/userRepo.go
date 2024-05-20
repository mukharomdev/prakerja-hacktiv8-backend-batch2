package repo

import (
	"gorm.io/gorm"
	"prakerja-session-7/models"
)


type UserRepo struct {
	UserDb *gorm.DB
}


func(ur *UserRepo)Create(usr *models.UserReq)(*models.User,error){
	newUser := models.User{
		Email	:usr.Email,
		Password:usr.Password,
	}
	err := ur.UserDb.Create(&newUser)

	if err.Error!=nil {
		return nil,err.Error
	}
	return &newUser,nil
}

func(ur *UserRepo)FindByEmail(usr *models.UserReq)(*models.User,error){
	userEmail := models.User{
		Email : usr.Email,
		Password:usr.Password,
	}

	err := ur.UserDb.First(&userEmail, "email = ?", userEmail.Email)
	if err.Error!=nil {
		return nil,err.Error
	}
	return &userEmail,nil
}

