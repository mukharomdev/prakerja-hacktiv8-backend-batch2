package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"net/http"
	"prakerja-sesi-5/dto"
)

var students []dto.Student = []dto.Student{}

func GenerateStudentId() int {
	if len(students) > 0 {
		studentId := students[len(students)-1].Id

		id, _ := strconv.Atoi(studentId)

		return id + 1
	}

	return 1
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	studentId := r.URL.Query().Get("id")

	index := 0

	for i, v := range students {
		if v.Id == studentId {
			index = i
			break
		}
	}

	if index == 0 {
		http.Error(w, "student not found", 404)
		return
	}

	students = append(students[:index], students[index+1:]...)

	response := dto.DeleteStudentResponse{
		Status: "success",
		Data:   nil,
	}

	b, err := json.Marshal(response)

	if err != nil {
		http.Error(w, "something went wrong", 500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(b)

}

func GetStudents(w http.ResponseWriter) {

	response := dto.GetStudentsResponse{
		Status: "success",
		Data:   students,
	}

	b, err := json.Marshal(response)

	_ = err

	w.Header().Add("Content-Type", "application/json")
	w.Write(b)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "invalid request body", 422)
		return
	}

	var requestBody dto.CreateStudentRequest

	err = json.Unmarshal(body, &requestBody)

	if err != nil {
		http.Error(w, "something went wrong", 500)
		return
	}

	student := dto.Student{
		Id:   fmt.Sprintf("%d", GenerateStudentId()),
		Name: requestBody.Name,
		Age:  requestBody.Age,
	}

	students = append(students, student)

	response := dto.CreateStudentResponse{
		Status: "success",
		Data:   student,
	}

	b, err := json.Marshal(response)

	if err != nil {
		http.Error(w, "something went wrong", 500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(201) // menuliskan status code terhadap response
	w.Write(b)
}

func main() {

	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {
			GetStudents(w)
			return
		}

		if r.Method == "POST" {
			CreateStudent(w, r)
			return
		}

		if r.Method == "DELETE" {
			DeleteStudent(w, r)
			return
		}
	})

	fmt.Println("Listening on PORT 8080")
	http.ListenAndServe(":8080", nil)
}
