package main

import (
	"api/internal/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	router.SetupRoutes()
	
	fmt.Println("Server Running in :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	
}