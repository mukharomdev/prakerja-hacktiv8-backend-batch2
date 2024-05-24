package service_test


import(
	"testing"
	"github.com/stretchr/testify/assert"
	//"errors"


	"uk-hacktiv8-prakerja/models"
	"uk-hacktiv8-prakerja/service"
	"uk-hacktiv8-prakerja/config"
	"uk-hacktiv8-prakerja/database"
	//repo "uk-hacktiv8-prakerja/repositories/postgres"

)

func TestUserServiceAdd(t *testing.T){
	config.Init()
	database.StartDatabase()

	var useres models.UserRegisterRes
	var usereq = models.UserRegisterReq{
		Username	: "Rishang",
		Email 		: "rishang@gmail.com",
		Password 	: "sukasuka",
		Age 		: 78,
	}



    var services service.UserService
	respons , err := services.Add(usereq)

 	useres = models.UserRegisterRes{
 		Username:respons.Username,
 	}
	assert.NoError(t,err)
	assert.Equal(t,respons.Username,useres.Username)

}


func TestUserServiceFindByUserName(t *testing.T){
	config.Init()
	database.StartDatabase()

	var userLoginres models.UserLoginRes
	var userLoginreq = models.UserLoginReq{
		Email 		: "yodhar@gmail.com",
		Password 	: "yodharishang",

	}



    var services service.UserService
	respons , err := services.FindByUserName(userLoginreq)

	err := bcrypt.CompareHashAndPassword([]byte(respons.Password), []byte(userLoginreq.Password))

 	userLoginres = models.UserLoginRes{
 		Password:respons.Password,
 	}
	assert.NoError(t,err)
	assert.Equal(t,respons.Password,userLoginres)

}