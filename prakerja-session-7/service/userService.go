package service

import (
	"prakerja-session-7/repo"
	"prakerja-session-7/models"
	//"strconv"
	//"log"
)

type UserService struct {
	UserRepo *repo.UserRepo
}

// func NewUserService(userRepo *repo.UserRepo)UserService{

//  return &UserServiceImpl{
//  	UserRepo : userRepo,
// 	}
// }

func(us *UserService)Store(user *models.UserReq)*models.UserResponse[models.User]{

   result,err := us.UserRepo.Save(user)

	if err!= nil {
		return &models.UserResponse[models.User]{
			Code:500,
			Success:nil,
			Error:&models.UserResponseError{
				Status: "error",
				Message: "internal server error",
			},
		}
	}
	return &models.UserResponse[models.User]{
		Code:201,
		Success:&models.UserResponseSuccess[models.User]{
			Status:"success",
			Data:*result,
		},
		Error:nil,
	}
}




func(us *UserService)FindByPassword(user *models.UserReq)*models.UserResponse[models.User]{
	users,err:= us.UserRepo.GetMany()


		if err!= nil {
		return &models.UserResponse[models.User]{
			Code:500,
			Success:nil,
			Error:&models.UserResponseError{
				Status: "error",
				Message: "internal server error",
			},
		}
	}

	var userPassword models.User

	for _,v := range users{
		if v.Email == user.Email{
			userPassword.Email = v.Email
			userPassword.ID = v.ID
		}
	}
	return &models.UserResponse[models.User]{
		Code:201,
		Success:&models.UserResponseSuccess[models.User]{
			Status:"success",
			Data:userPassword,
		},
		Error:nil,
	}


}
