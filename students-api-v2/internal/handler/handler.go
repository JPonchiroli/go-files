package handler

import (
	"fmt"
	"net/http"
	"encoding/json"
	"api/internal/model"
	"api/internal/service"
	"github.com/go-chi/chi/v5"
)

type StudentHandler struct {
	Service *service.StudentService
}

func NewStudentHandler(service *service.StudentService) * StudentHandler {

	return &StudentHandler{
		Service: service,
	}

}

func TestConnection(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprint(w, "Hello, GET!")

}

func (h *StudentHandler) GetStudents(w http.ResponseWriter, r *http.Request,) {

}

func PostStudent(w http.ResponseWriter, r *http.Request) {

	var students model.StudentRequest

	err := json.NewDecoder(r.Body).Decode(&students)

	if err != nil {

		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return

	}

	service.PostStudent(&students)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(students)

}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {

	studentName := chi.URLParam(r, "name")
	
	err := json.NewDecoder(r.Body).Decode(&studentName)

	if err != nil {

		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return

	}

	service.DeleteStudent(studentName)

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(true)

}