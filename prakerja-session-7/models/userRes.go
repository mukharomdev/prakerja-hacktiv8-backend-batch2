package models


type UserListOrOne interface {
	[]User | User
}

type UserResponseSuccess[T UserListOrOne] struct {
	Status string `json:"status"`
	Data T `json:"data"`
}

type UserResponseError struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

type UserResponse[T UserListOrOne] struct{
	Code int
	Success *UserResponseSuccess[T]
	Error *UserResponseError

}