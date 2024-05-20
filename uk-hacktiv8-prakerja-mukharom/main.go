package main

import (
	"uk-hacktiv8-prakerja-mukharom/config"
	"uk-hacktiv8-prakerja-mukharom/database"
	"uk-hacktiv8-prakerja-mukharom/server"
)

func main() {
	config.Init()

	database.StartDatabase()
	s := server.NewServer()

	s.Run()
}
