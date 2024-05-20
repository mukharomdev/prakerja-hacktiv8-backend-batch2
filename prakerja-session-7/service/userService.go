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

func(us *UserService)Create(user *models.UserReq)*models.UserResponse[models.User]{

   result,err := us.UserRepo.Create(user)

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




func(us *UserService)FindByEmail(user *models.UserReq)*models.UserResponse[models.User]{
	result,err := us.UserRepo.FindByEmail(user)

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
