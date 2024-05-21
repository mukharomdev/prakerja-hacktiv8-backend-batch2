package testpostgres


import(
	"fmt"
	"testify"
	"testing"
	"github.com/stretchr/testify/assert"
	"uk-hacktiv8-prakerja/config"
	"uk-hacktiv8-prakerja/models"
	"uk-hacktiv8-prakerja/database"
	repo "uk-hacktiv8-prakerja/repositories/postgres"
)

func TestUserRepo(t *testing.T){
	config.Init()
	database.StartDatabase()
	// instaniasi model user
	var user models.User
	// instaniasi struct userRepo
	repos := &repo.UserRepository{}
    result,err = repos.Store(&user)
	// Mulai test
	assert.Equal(t,result,user,"harus berupa objek user")

}