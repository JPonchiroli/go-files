package router

import (
	"net/http"
	"api/internal/handler"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes() http.Handler {

	r := chi.NewRouter()
	
	http.HandleFunc("/get", handler.TestConnection)
	
	r.Get("/students", handler.GetStudent)
	r.Post("/students", handler.PostStudent)
	r.Delete("/students/{id}", handler.DeleteStudent)

	return r
}