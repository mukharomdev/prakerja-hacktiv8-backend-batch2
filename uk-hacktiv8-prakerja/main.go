package main

import (
	//"fmt"
	"github.com/gin-gonic/gin"


	"uk-hacktiv8-prakerja/config"
	//"uk-hacktiv8-prakerja/models"
	"uk-hacktiv8-prakerja/database"
	"uk-hacktiv8-prakerja/routers"
	//repo "uk-hacktiv8-prakerja/repositories/postgres"
)

func main(){
	config.Init()

	database.StartDatabase()

	router := gin.Default()


	server := routers.Routers(router)


	server.Run(":8000")







}