package dto

type Student struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type CreateStudentRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type CreateStudentResponse struct {
	Status string  `json:"status"`
	Data   Student `json:"data"`
}

type GetStudentsResponse struct {
	Status string    `json:"status"`
	Data   []Student `json:"data"`
}

type DeleteStudentResponse struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}
