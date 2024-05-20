package main

// import (
// 	"prakerja-session-6/models"
// 	//"prakerja-session-6/controllers"
// 	"prakerja-session-6/routers"
// 	//"github.com/gin-gonic/gin"

// )

import (
	"github.com/gin-gonic/gin"
	 "gorm.io/gorm"
	 "gorm.io/driver/postgres"
	 //"os"
	 "fmt"
	 "log"
	 "prakerja-session-6/models"
	 "prakerja-session-6/handler"
	 "prakerja-session-6/repo"
	 "prakerja-session-6/service"
)


var (
	host     = "localhost"
	port     = "5432"
	user     = "yodharishang"
	password = "yodha3129"
	dbName   = "product_db"

)

func main(){
	// router engine gin
	ServerGin := gin.Default()

	// koneksi ke postgresql
	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil{
		log.Fatal("error conecting to Database",err)
	}

	//Database migrate
	db.AutoMigrate(&models.Product{})

	// product repo
	productRepo := &repo.ProductRepo{DB : db}

	// product service

	productService := &service.ProductService{ProductRepo:productRepo}

	// Routing
	Router := handler.ProductsHandler{
		EngineGin:ServerGin,
		Service:productService,
	}

    // Handling  server route
	Router.Route()

	// menjalankan server
	PORT := ":8001"
    ServerGin.Run(PORT)

}