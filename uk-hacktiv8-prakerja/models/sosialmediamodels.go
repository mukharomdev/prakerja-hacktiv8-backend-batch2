package models

type SocialMedia struct{
	ID 			uint 	  	`gorm:"primary_key" json:"id,omitempty"`
	Name 		string 		`gorm:"not null" json:"name,omitempty"`
	Social_Media_Url string `gorm:"not null" json:"social_media_url,omitempty"`
}
