package service


import(
	"uk-hacktiv8-prakerja/models"
	repo "uk-hacktiv8-prakerja/repositories/postgres"
)


type UserService struct{
	Repo  *repo.UserRepository
}

func(s *UserService)Add(usreq models.UserRegisterReq )(*models.UserRegisterRes,error){
	// user := models.User{
	// 	Username :usreq.Username,
	// 	Email	 :usreq.Email,
	// 	Password :usreq.Password,
	// 	Age 	 :usreq.Age,
	// }

	user := models.NewUser(&usreq)

	respons , err := s.Repo.Store(*user)

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
			userloginres.Password = v.Password
		}
	}

	return &userloginres,nil
}