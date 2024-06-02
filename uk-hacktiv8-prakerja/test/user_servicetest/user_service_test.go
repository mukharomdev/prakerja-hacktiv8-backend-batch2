package service_test


import(
	"testing"
	"github.com/stretchr/testify/assert"
	//"errors"
    "golang.org/x/crypto/bcrypt"

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

	//var userLoginres models.UserLoginRes
	var userLoginreq = models.UserLoginReq{
		Email 		: "yodhar@gmail.com",
		Password 	: "yodharishang",

	}



    var services service.UserService
	respons , err := services.FindByUserName(userLoginreq)

	validity := bcrypt.CompareHashAndPassword([]byte(respons.Password), []byte(userLoginreq.Password))

 	// userLoginres = models.UserLoginRes{
 	// 	Password:respons.Password,
 	// }
 	assert.NoError(t,validity)
	assert.Error(t,err)
	assert.Equal(t,respons.Password,userLoginreq.Password)

}



func TestUserServiceUpdate(t *testing.T){
	config.Init()
	database.StartDatabase()

	//var userLoginres models.UserLoginRes
	var userUpdateReq = models.UserUpdateReq{
		Email 		: "mukharomdev@gmail.com",
		Username 	: "mukharomdev",

	}



    var services service.UserService
	respons , err := services.Update(userUpdateReq,1)
    t.Log(respons)

    assert.Error(t,err)


}