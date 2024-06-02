package postgres


import (
	"fmt"
	"uk-hacktiv8-prakerja/database"
	"uk-hacktiv8-prakerja/models"
)


type UserRepository struct{

}

func(u *UserRepository)Store(usr models.User)(models.User,error){
	db := database.GetDatabase()
	err := db.Create(&usr).Error

	if err != nil {
		return models.User{}, err
	}

	return usr, nil
}

func (s *UserRepository) Delete(id uint) (models.User, error) {
	db := database.GetDatabase()
	var n models.User
	if !s.IsExist(id){
		return models.User{},nil
	}else{
		n, err := s.Get(id)
		if err != nil {
			return models.User{}, err
		}

		err = db.Delete(&n,id).Error
		if err != nil {
			return models.User{}, fmt.Errorf("tidak bisa menghapus data: %v", err)
		}

	}


	return n, nil
}

func (s *UserRepository) Update(p models.User) error {
	db := database.GetDatabase()
	//err := db.Save(&p).Error
	err := db.Model(&p).Where("id = ?",p.ID).Updates(&p)
		if err != nil {
			return fmt.Errorf("tidak bisa mengupdate user di database: %v", err)
		}



	return nil
}

func (s *UserRepository) Get(id uint) (*models.User, error) {
	db := database.GetDatabase()
	var user models.User

	//err := db.First(&user, "id = ?", id).Error
	err := db.Model(&user).First(&user, id).Error

	if err != nil {
		return &models.User{},err
	}

	return &user, nil
}

func (s *UserRepository) GetAll() ([]models.User, error) {
	db := database.GetDatabase()
	var p []models.User
	err := db.Find(&p).Error

	if err != nil {
		return nil, fmt.Errorf("cannot find Users: %v", err)
	}

	return p, err
}

func (s *UserRepository)IsExist(id uint)(bool){
	db := database.GetDatabase()
	user := &models.User{
		ID:id,
	}

	 result := db.First(user)

	 return result.RowsAffected > 0

}

