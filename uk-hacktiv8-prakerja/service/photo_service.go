package service


import(
	//"log"
	//"fmt"
	"uk-hacktiv8-prakerja/models"
	repo "uk-hacktiv8-prakerja/repositories/postgres"
)


type PhotoService struct{
	Repo  *repo.PhotoRepository
	Usr   *repo.UserRepository
}

func(s *PhotoService)Add(preq models.PhotoCreateReq )(*models.PhotoCreateRes,error){
	Photo := models.Photo{
		Title 	 :preq.Title,
		Caption	 :preq.Caption,
		Photo_url:preq.Photo_url,
		UserID	 :preq.UserID,
	}


    // save data photo ke database
	respons , err := s.Repo.Store(Photo)

	if err != nil {
		return &models.PhotoCreateRes{},err
	}
	var photocreateRes models.PhotoCreateRes

	photocreateRes = models.PhotoCreateRes{
			ID 			: respons.ID,
			Title 		: respons.Title,
			Caption 	: respons.Caption,
			Photo_url 	: respons.Photo_url,
			UserID		: respons.UserID,
			Created_at  : respons.Created_at,
	}

	return &photocreateRes,nil
}

func(s *PhotoService)FindById(usrid uint)([]models.PhotoGetRes,error){

	// Photos,err := s.Repo.GetAll()
	usr,err := s.Usr.Get(usrid)

	if err != nil {
		return []models.PhotoGetRes{},nil
	}


	//log.Printf("%v",usr)



	photos,err := s.Repo.GetAll()

    if err != nil {
		return []models.PhotoGetRes{},nil
	}

    //log.Printf("%v",photos)

    var usrphoto models.UserPhoto

    usrphoto = models.UserPhoto{
    	Email 	 : usr.Email,
    	Username : usr.Username,
    }

    var photoGetRes models.PhotoGetRes
    var photoAllGetRes []models.PhotoGetRes


	for i,_ := range photos{

    	photoGetRes = models.PhotoGetRes{
    	Photo  :photos[i],
    	UserPh :usrphoto,
    	}

    	photoAllGetRes = append(photoAllGetRes,photoGetRes)
    }




	return photoAllGetRes,nil
}

// func(s *PhotoService)Update(Photoeq models.PhotoUpdateReq,id uint)(*models.PhotoUpdateRes,error){
// 	PhotoUpdate := models.Photo{
// 		ID 		: id,
// 		Email 	: Photoeq.Email,
// 		Photoname: Photoeq.Photoname,
// 		}
// 	s.Repo.Update(PhotoUpdate)

// 	respons,err := s.Repo.Get(PhotoUpdate.ID)

// 	if err != nil{
// 		return &models.PhotoUpdateRes{},nil
// 		}


// 	updateRespons := &models.PhotoUpdateRes{
// 		ID 		  : respons.ID,
// 		Photoname  : respons.Photoname,
// 		Email 	  : respons.Email,
// 		Age 	  : respons.Age,
// 		Updated_at: respons.Updated_at,
// 		}





// 	return updateRespons,nil

// }


// func(s *PhotoService)Delete(id uint)(error){
// 	PhotoDelete := models.Photo{ID : id }
// 	respons,err := s.Repo.Delete(PhotoDelete.ID)

// 	_ = respons

// 	// if err != nil{
// 	// 	return &models.Photo{},nil
// 	// }
// 	return err

// }

