package dto

type LoginRequestDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponseDto struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}
