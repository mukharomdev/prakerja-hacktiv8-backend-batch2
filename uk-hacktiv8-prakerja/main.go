package main

import (
	"fmt"
	"uk-hacktiv8-prakerja/config"
	"uk-hacktiv8-prakerja/models"
	"uk-hacktiv8-prakerja/database"
	repo "uk-hacktiv8-prakerja/repositories/postgres"
)

func main(){
	config.Init()

	database.StartDatabase()

	var user *models.User

	user = &models.User{
		Email:"pelatge@gmail.com",
		Password:"yodharishang",
	}

	r := &repo.UserRepository{}
	v,err := r.Store(*user)

	if err != nil{
		fmt.Printf("%v===%v",v,err)
	}



}