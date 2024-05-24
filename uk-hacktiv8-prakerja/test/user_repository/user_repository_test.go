package postgres_test


import(
	//"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	//"github.com/rubenv/pgtest"

	"uk-hacktiv8-prakerja/config"
	"uk-hacktiv8-prakerja/models"
	"uk-hacktiv8-prakerja/database"
	repo "uk-hacktiv8-prakerja/repositories/postgres"
)

func TestUserRepositoryStore(t *testing.T){
	config.Init()
	database.StartDatabase()
 	// Do something with pg.DB (which is a *sql.DB)
   	// instaniasi model user
	var user models.User
	user = models.User{
			Username :"mukharoms",
			Email	 :"yodharishangs@gmail",
			Password :"golangindonesias",
			Age		 : 38,
	}
	// instaniasi struct userRepo
			repos := &repo.UserRepository{}
		    result,err:= repos.Store(user)
		   // t.Log(result)
    // mulai test

    		assert.NoError(t,err)
    		assert.Equal(t,result.Username,user.Username)

}

func TestUserRepositoryDelete(t *testing.T){
	config.Init()
	database.StartDatabase()

	// instaniasi model user
	//var user models.User
	// instaniasi struct userRepo
	repos := &repo.UserRepository{}
    result,err:= repos.Delete(1)
    //t.Log(result)

    	// Mulai test
    assert.NoError(t, err)
	assert.Equal(t,result,models.User{},"objek user harus models.User{}")



}

func TestUserRepositoryUpdate(t *testing.T){
	config.Init()
	database.StartDatabase()
	// instaniasi model user
	var user models.User
	// instaniasi struct userRepo
	repos := &repo.UserRepository{}
    err   := repos.Update(user)

    	// Mulai test
    assert.NoError(t, err)
    //assert.Equal(t,result,models.User{},"objek user harus models.User{}")



}

func TestUserRepositoryGet(t *testing.T){
	config.Init()
	database.StartDatabase()
	// instaniasi model user
	var user models.User
	// instaniasi struct userRepo
	repos := &repo.UserRepository{}
    result,err:= repos.Get(1)
    //t.Log(result)

    	// Mulai test
    assert.NoError(t, err)
    assert.Equal(t,result.ID,user.ID,"harus berupa objek user")


}

func TestUserRepositoryGetAll(t *testing.T){
	config.Init()
	database.StartDatabase()
	// instaniasi model user
	var user models.User
	user = models.User{
			Username :"mukharom",
			Email	 :"yodharishang@gmail",
			Password :"golangindonesia",
			Age		 : 38,
	}
	// instaniasi struct userRepo
			repos := &repo.UserRepository{}
		    repos.Store(user)
		    //t.Log(result1)
    // mulai
	// instaniasi struct userRepo
	//repos = &repo.UserRepository{}
    result,err := repos.GetAll()
    //t.Log(result)
    result2,_ := repos.GetAll()

    	// Mulai test
    assert.NoError(t, err)
    assert.Equal(t,result,result2,"harus berupa slice objek user")


}

