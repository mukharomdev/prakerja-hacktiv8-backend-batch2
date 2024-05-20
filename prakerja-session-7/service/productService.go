package service

import (
	"prakerja-session-7/repo"
	"prakerja-session-7/models"
	"strconv"
	//"log"
)

type ProductService struct {
	ProductRepo *repo.ProductRepo
}

func (ps *ProductService) GetAll()*models.ProdResponse[[]models.Product]{
	result,err:= ps.ProductRepo.GetMany()
	if err!=nil {
		//log.Fatalf(err.Error())
		return &models.ProdResponse[[]models.Product]{
			Code:500,
			Success:nil,
			Error:&models.ProdResponseError{
				Status:"error",
				Message:"internal server error",
			},
		}
	}
	return &models.ProdResponse[[]models.Product]{
		Code:200,
		Success:&models.ProdResponseSuccess[[]models.Product]{
			Status:"success",
			Data:result,
		},
		Error:nil,
	}
}

func (ps *ProductService) Add(product *models.ProductReq)*models.ProdResponse[models.Product]{
	result,err:=ps.ProductRepo.Save(product)
	if err!= nil {
		return &models.ProdResponse[models.Product]{
			Code:500,
			Success:nil,
			Error:&models.ProdResponseError{
				Status: "error",
				Message: "internal server error",
			},
		}
	}
	return &models.ProdResponse[models.Product]{
		Code:201,
		Success:&models.ProdResponseSuccess[models.Product]{
			Status:"success",
			Data:*result,
		},
		Error:nil,
	}
}

func (ps *ProductService) Edit(id string,product *models.ProductReq)*models.ProdResponse[models.Product]{
	iid,err:=strconv.Atoi(id)
	if err!=nil {
		return &models.ProdResponse[models.Product]{
			Code:400,
			Success:nil,
			Error:&models.ProdResponseError{
				Status:"fail",
				Message:"bad request",
			},
		}
	}
	if !ps.ProductRepo.IsExist(uint(iid)) {
		return &models.ProdResponse[models.Product]{
			Code:404,
			Success:nil,
			Error:&models.ProdResponseError{
				Status:"fail",
				Message:"not found",
			},
		}
	}

	result,err:=ps.ProductRepo.Edit(uint(iid),product)
	if err!= nil {
		//log.Fatalf(err.Error())
		return &models.ProdResponse[models.Product]{
			Code:500,
			Success:nil,
			Error:&models.ProdResponseError{
				Status:"error",
				Message:"internal server error",
			},
		}
	}
	return &models.ProdResponse[models.Product]{
		Code:200,
		Success:&models.ProdResponseSuccess[models.Product]{
			Status:"success",
			Data:*result,
		},
		Error:nil,
	}
}

func (ps *ProductService) Delete(id string)*models.ProdResponse[[]models.Product]{
	iid,err:=strconv.Atoi(id)
        if err!=nil {
                return &models.ProdResponse[[]models.Product]{
                        Code:400,
                        Success:nil,
                        Error:&models.ProdResponseError{
                                Status:"fail",
                                Message:"bad request",
                        },
                }
        }
        if !ps.ProductRepo.IsExist(uint(iid)) {
                return &models.ProdResponse[[]models.Product]{
                        Code:404,
                        Success:nil,
                        Error:&models.ProdResponseError{
                                Status:"fail",
                                Message:"not found",
                        },
                }
        }

        err=ps.ProductRepo.Delete(uint(iid))
        if err!= nil {
                //log.Fatalf(err.Error())
                return &models.ProdResponse[[]models.Product]{
                        Code:500,
                        Success:nil,
                        Error:&models.ProdResponseError{
                                Status:"error",
                                Message:"internal server error",
                        },
                }
        }
        return &models.ProdResponse[[]models.Product]{
                Code:200,
                Success:&models.ProdResponseSuccess[[]models.Product]{
                        Status:"success",
                        Data:nil,
                },
                Error:nil,
        }
}





