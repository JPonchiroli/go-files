package router

import (
	"net/http"
	"api/internal/handler"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes() {
	http.HandleFunc("/get", handler.TestConnection)

	//r := chi.NewRouter() Define a switch to decide which method use
}