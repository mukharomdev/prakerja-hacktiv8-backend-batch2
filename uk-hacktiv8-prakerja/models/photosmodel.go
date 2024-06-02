package models

import (
	"time"

	//"github.com/google/uuid"
)




type Photo struct{
	ID 			uint 	  `gorm:"primary_key" json:"id,omitempty"`
	Title 		string 	  `gorm:"not null" json:"title,omitempty"`
	Caption 	string 	  `json:"caption,omitempty"`
	Photo_url 	string 	  `gorm:"not null" json:"photo_url,omitempty"`
	UserID		uint
	Created_at  time.Time  `gorm:"autoCreateTime" json:"created_at,omitempty"`
	Updated_at  time.Time  `gorm:"autoUpdateTime " json:"updated_at,omitempty"`
}


// REQUEST
// ● headers: Authorization (Bearer token string)
// POST /photos
type PhotoPostReq struct{
	Title 		string 	  `json:"title"`
	Caption 	string 	  `json:"caption"`
	Photo_url 	string 	  `json:"photo_url"`
}
// Res
type PhotoRes struct{
	ID 			uint 	  `json:"id,omitempty"`
	Title 		string 	  `json:"title,omitempty"`
	Caption 	string 	  `json:"caption,omitempty"`
	Photo_url 	string 	  `json:"photo_url,omitempty"`
	UserID		uint
	Created_at  time.Time  `json:"created_at,omitempty"`

}


// GET /photos
// ● headers: Authorization (Bearer token string)

//res
type PhotoListRes struct{
	Photos []Photo
}

// PUT /photos/:photoId
// Request
// ● headers: Authorization (Bearer token string)
// ● params: photoId (string)
// ● body:

type PhotosPutReq struct {

}
