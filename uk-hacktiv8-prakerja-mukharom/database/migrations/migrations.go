package migrations

import (
	"uk-hacktiv8-prakerja-mukharom/models"
	"gorm.io/gorm"
)

func RunAutoMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.User{},&models.Comment{},&models.Photo{},models.SocialMedia{})
}
