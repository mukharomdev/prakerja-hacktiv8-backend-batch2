package models
import (
	"time"

	//"github.com/google/uuid"
)

type Comment struct {
	ID 			uint 	  `gorm:"primary_key" json:"id,omitempty"`
	UserID 		uint 	  `json:"user_id,omitempty"`
	Photo_ID 	uint 	  `json:"photo_id,omitempty"`
	Message 	string 	  `gorm:"not null" json:"message,omitempty"`
	Created_at  time.Time  `gorm:"autoCreateTime" json:"created_at,omitempty"`
	Updated_at  time.Time  `gorm:"autoUpdateTime " json:"updated_at,omitempty"`
}
