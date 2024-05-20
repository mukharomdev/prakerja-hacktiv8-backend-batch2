package test

import (
	"testing"
	//"net/http"
	"uk-hacktiv8-prakerja-mukharom/models"
	"uk-hacktiv8-prakerja-mukharom/controllers"
  	//"io"
	//"encoding/json"
	//"fmt"
	//"bytes"
	//"gorm.io/gorm"
	//"github.com/gin-gonic/gin"

)


func TestAPI(t *testing.T){
 var ac controllers.AuthUserController
 var user models.User
    err := ac.DB.Create(&user)
    if err != nil {
    	t.Log(err)
    }




}
