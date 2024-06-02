package migrations

import (
	"uk-hacktiv8-prakerja/models"
	"gorm.io/gorm"
)

func RunAutoMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.User{},&models.Photo{},&models.SocialMedia{},&models.Comment{})
}
