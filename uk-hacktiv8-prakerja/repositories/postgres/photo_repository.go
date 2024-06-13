package postgres


import (
	"fmt"
	"uk-hacktiv8-prakerja/database"
	"uk-hacktiv8-prakerja/models"
)


type PhotoRepository struct{

}

func(u *PhotoRepository)Store(p models.Photo)(models.Photo,error){
	db := database.GetDatabase()
	err := db.Create(&p).Error

	if err != nil {
		return models.Photo{}, err
	}

	return p, nil
}

// func (s *PhotoRepository) Delete(id uint) (models.Photo, error) {
// 	db := database.GetDatabase()
// 	var n models.Photo
// 	if !s.IsExist(id){
// 		return models.Photo{},nil
// 	}else{
// 		n, err := s.Get(id)
// 		if err != nil {
// 			return models.Photo{}, err
// 		}

// 		err = db.Delete(&n,id).Error
// 		if err != nil {
// 			return models.Photo{}, fmt.Errorf("tidak bisa menghapus data: %v", err)
// 		}

// 	}


// 	return n, nil
// }

// func (s *PhotoRepository) Update(p models.Photo) error {
// 	db := database.GetDatabase()
// 	//err := db.Save(&p).Error
// 	err := db.Model(&p).Where("id = ?",p.ID).Updates(&p)
// 		if err != nil {
// 			return fmt.Errorf("tidak bisa mengupdate Photo di database: %v", err)
// 		}



// 	return nil
// }

func (s *PhotoRepository) Get(userid uint) (*models.Photo, error) {
	db := database.GetDatabase()
	var Photo models.Photo

	//err := db.First(&Photo, "id = ?", id).Error
	err := db.Model(&Photo).First(&Photo,"user_id = ?", userid).Error

	if err != nil {
		return &models.Photo{},err
	}

	return &Photo, nil
}

func (s *PhotoRepository) GetAll() ([]models.Photo, error) {
	db := database.GetDatabase()
	var p []models.Photo
	err := db.Find(&p).Error

	if err != nil {
		return nil, fmt.Errorf("cannot find Photos: %v", err)
	}

	return p, err
}

// func (s *PhotoRepository)IsExist(id uint)(bool){
// 	db := database.GetDatabase()
// 	Photo := &models.Photo{
// 		ID:id,
// 	}

// 	 result := db.First(Photo)

// 	 return result.RowsAffected > 0

// }

