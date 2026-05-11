package main

import (
	"api/internal/router"
	"log"
	"net/http"
)

func main() {
	r := router.SetupRoutes()
	
	log.Println("Server Running in :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}