package models

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRes struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}
