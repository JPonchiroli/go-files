package handler

import (

	"github.com/go-chi/chi/v5"

)

func SetupRoutes(studentHandler *StudentHandler) *chi.Mux {

	r := chi.NewRouter()

	r.Get("/students", studentHandler.GetStudents)

	return r
	
}