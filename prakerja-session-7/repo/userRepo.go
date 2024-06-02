package repo

import (
	"gorm.io/gorm"
	"prakerja-session-7/models"
)


type UserRepo struct {
	DB *gorm.DB
}


func (usr *UserRepo) IsExist(id uint)bool{
	User:= models.User{ID:id}

	result:=usr.DB.Find(&User)

	return result.RowsAffected > 0
}

func (usr *UserRepo) GetMany()([]models.User,error){
		var Users []models.User

		result := usr.DB.Find(&Users)

		if result.Error!=nil {

			return nil,result.Error
		}
		return Users,nil
}


func (usr *UserRepo) Save(usrodReq *models.UserReq)(*models.User,error){
	User := models.User{
		Email: usrodReq.Email,
		Password: usrodReq.Password,
	}

	result := usr.DB.Create(&User)

	if result.Error!=nil {
		return nil,result.Error
	}
	return &User,nil
}

func (usr *UserRepo) Edit(id uint,UserReq *models.UserReq)(*models.User,error){
	User:= models.User{
		ID:id,
	}
	result:=usr.DB.Model(&User).Updates(UserReq)

	if result.Error!=nil {

		return nil,result.Error
	}
	return &User,nil


}


func (usr *UserRepo) Delete(id uint)error{

	result := usr.DB.Delete(&models.User{ID:id})

	return result.Error
}



