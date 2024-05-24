package interfaces

import(
	"uk-hacktiv8-prakerja/models"
)


type IUserService interface{
	Add(usreq models.UserLoginReq)(usres models.UserLoginRes,err error)
	FindByUserName(usreq models.UserLoginReq)(*models.UserLoginRes,error)
}