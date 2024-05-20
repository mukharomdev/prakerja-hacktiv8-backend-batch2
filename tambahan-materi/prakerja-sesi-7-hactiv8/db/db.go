package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (

	timeZone = "Asia/Jakarta"
	sslMode  = "disable"

	host     = "localhost"
	port     = "5432"
	user     = "yodharishang"
	password = "yodha3129"
	dbname   = "product_db"

)

var (
	db  *gorm.DB
	err error
)

func InitializeDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslMode, timeZone,
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	//db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("errror connecting to db: (%s)", err.Error())
	}
}

func GetDB() *gorm.DB {
	return db
}
