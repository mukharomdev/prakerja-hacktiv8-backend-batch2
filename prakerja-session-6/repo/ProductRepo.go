package repo

import (
	"gorm.io/gorm"
	"prakerja-session-6/models"
)


type ProductRepo struct {
	DB *gorm.DB
}


func (pr *ProductRepo) IsExist(id uint)bool{
	product:= models.Product{ID:id}

	result:=pr.DB.Find(&product)

	return result.RowsAffected > 0
}

func (pr *ProductRepo) GetMany()([]models.Product,error){
		var products []models.Product

		result := pr.DB.Find(&products)

		if result.Error!=nil {

			return nil,result.Error
		}
		return products,nil
}


func (pr *ProductRepo) Save(prodReq *models.ProductReq)(*models.Product,error){
	product := models.Product{
		Name: prodReq.Name,
		Price: prodReq.Price,
	}

	result:=pr.DB.Create(&product)

	if result.Error!=nil {
		return nil,result.Error
	}
	return &product,nil
}

func (pr *ProductRepo) Edit(id uint,productReq *models.ProductReq)(*models.Product,error){
	product:= models.Product{
		ID:id,
	}
	result:=pr.DB.Model(&product).Updates(productReq)

	if result.Error!=nil {

		return nil,result.Error
	}
	return &product,nil


}


func (pr *ProductRepo) Delete(id uint)error{

	result := pr.DB.Delete(&models.Product{ID:id})

	return result.Error
}

