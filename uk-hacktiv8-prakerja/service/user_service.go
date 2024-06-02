package service


import(
	//"log"
	//"fmt"
	"uk-hacktiv8-prakerja/models"
	repo "uk-hacktiv8-prakerja/repositories/postgres"
)


type UserService struct{
	Repo  *repo.UserRepository
}

func(s *UserService)Add(usreq models.UserRegisterReq )(*models.UserRegisterRes,error){
	user := models.User{
		Username :usreq.Username,
		Email	 :usreq.Email,
		Password :usreq.Password,
		Age 	 :usreq.Age,
	}

	// user := models.User(&usreq)

	respons , err := s.Repo.Store(user)

	if err != nil {
		return &models.UserRegisterRes{},err
	}
	var useregisterres models.UserRegisterRes
	useregisterres = models.UserRegisterRes{
		Age 	:respons.Age,
		Email	:respons.Email,
		ID 	 	:respons.ID,
		Username:respons.Username,
	}

	return &useregisterres,nil
}

func(s *UserService)FindByUserName(usreq models.UserLoginReq)(*models.UserLoginRes,error){

	var userloginres models.UserLoginRes

	users,err := s.Repo.GetAll()

	if err != nil {
		return &models.UserLoginRes{},nil
	}

	for _,v := range users{
		if v.Email == usreq.Email{
			userloginres.Email 	  = v.Email
			userloginres.Password = v.Password
			userloginres.ID		  = v.ID
		}
	}

	return &userloginres,nil
}

func(s *UserService)Update(usereq models.UserUpdateReq,id uint)(*models.UserUpdateRes,error){
	userUpdate := models.User{
		ID 		: id,
		Email 	: usereq.Email,
		Username: usereq.Username,
		}
	s.Repo.Update(userUpdate)

	respons,err := s.Repo.Get(userUpdate.ID)

	if err != nil{
		return &models.UserUpdateRes{},nil
		}


	updateRespons := &models.UserUpdateRes{
		ID 		  : respons.ID,
		Username  : respons.Username,
		Email 	  : respons.Email,
		Age 	  : respons.Age,
		Updated_at: respons.Updated_at,
		}





	return updateRespons,nil

}


func(s *UserService)Delete(id uint)(error){
	userDelete := models.User{ID : id }
	respons,err := s.Repo.Delete(userDelete.ID)

	_ = respons

	// if err != nil{
	// 	return &models.User{},nil
	// }
	return err

}

