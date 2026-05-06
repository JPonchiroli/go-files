package handler

import (
	"api/internal/service"
	"encoding/json"
	"fmt"
	"net/http"
)

func TestConnection(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprint(w, "Hello, GET!")

}

func GetStudent(w http.ResponseWriter, r *http.Request) {

	student := service.GetStudent()
	json.NewEncoder(w).Encode(student)

}

func GetStudentById(w http.ResponseWriter, r *http.Request) {

}

func PostStudent(w http.ResponseWriter, r *http.Request) {

}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {

}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {

}