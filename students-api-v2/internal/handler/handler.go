package handler

import (
	"fmt"
	"net/http"
	"encoding/json"
	"api/internal/model"
	"api/internal/service"
	"github.com/go-chi/chi/v5"
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

	var students model.StudentRequest

	verifyErr(w, r, &students)

	service.PostStudent(&students)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(students)

}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {

	studentName := chi.URLParam(r, "name")
	
	verifyErr(w, r, studentName)

	service.DeleteStudent(studentName)

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(true)

}

func verifyErr(w http.ResponseWriter, r *http.Request, element any){

	err := json.NewDecoder(r.Body).Decode(&element)

	if err != nil {

		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return

	}

}