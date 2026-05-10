package handler

import (
	"api/internal/model"
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

func PostStudent(w http.ResponseWriter, r *http.Request) {

	var student model.Student

	err := json.NewDecoder(r.Body).Decode(&student)

	if err != nil {

		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return

	}

	service.PostStudent(&student)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(students)

}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {

}