package router

import (
	"net/http"
	"api/internal/handler"
)

func SetupRoutes() {
	http.HandleFunc("/get", handler.TestConnection)
}