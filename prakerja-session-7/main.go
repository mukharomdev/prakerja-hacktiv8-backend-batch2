package main



import (
	"github.com/gin-gonic/gin"
	 "gorm.io/gorm"
	 "gorm.io/driver/postgres"
	 //"os"
	 "fmt"
	 "log"
	 "prakerja-session-7/models"
	 "prakerja-session-7/handler"
	 "prakerja-session-7/repo"
	 "prakerja-session-7/service"
	 "prakerja-session-7/controllers"
)


var (
	host     = "localhost"
	port     = "5432"
	user     = "yodharishang"
	password = "yodha3129"
	dbName   = "product_db"

)

func main(){
	fmt.Println("Belajar Golang Middleware  prakerja-Hactiv8")
	// router engine gin
	ServerGin := gin.Default()

	// koneksi ke postgresql
	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil{
		log.Fatal("error conecting to Database",err)
	}

	//Database migrate
	db.AutoMigrate(&models.Product{},&models.User{})

	// product repo
	productRepo := &repo.ProductRepo{DB : db}

	// product service

	productService := &service.ProductService{ProductRepo:productRepo}

	// Routing
	ProductHandler := handler.ProductsHandler{
		EngineGin:ServerGin,
		Service:productService,
	}

    // Handling  server route
	ProductHandler.Route()
    /*==========================================================================
      ==========================================================================
    */
    //User repo
    db.AutoMigrate(&models.User{})

	userRepo := &repo.UserRepo{DB : db}

	// product service

	userService := &service.UserService{UserRepo:userRepo}

	// Routing
	UserHandler:= controllers.UserController{
		Engine:ServerGin,
		UserService:userService,
	}

    // Handling  server route
	UserHandler.Route()





	// menjalankan server
	PORT := ":8001"
    ServerGin.Run(PORT)

}