package models

// type UserReq struct {
// 	Status string `json:"status"`
// 	Data   string `json:"data"`
// }

type UserReqAuth struct{
	Login string
	Register string
}

type UserReq struct{
	Email string 	`json:"email"`
	Password string `json:"password"`
}

type UserDataRegister struct {
	ID uint `json:"id"`
	Email string `json:"email"`
}