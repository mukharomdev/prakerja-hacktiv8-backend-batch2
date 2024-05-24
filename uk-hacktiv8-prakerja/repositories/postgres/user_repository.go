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

		err = db.Delete(&n).Error
		if err != nil {
			return models.User{}, fmt.Errorf("tidak bisa menghapus data: %v", err)
		}

	}


	return n, nil
}

func (s *UserRepository) Update(p models.User) error {
	db := database.GetDatabase()
	if s.IsExist(p.ID){
		err := db.Save(&p).Error
		if err != nil {
			return fmt.Errorf("tidak bisa mengupdate user di database: %v", err)
		}
	}


	return nil
}

func (s *UserRepository) Get(id uint) (models.User, error) {
	db := database.GetDatabase()
	var p models.User
	err := db.First(&p, id).Error

	if err != nil {
		return models.User{}, fmt.Errorf("tidak dapat menemukan user dengan  id: %v", err)
	}

	return p, nil
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

