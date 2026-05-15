package main

import (
	"api/internal/database"
	"api/internal/handler"
	"api/internal/repository"
	"api/internal/service"
	"log"
	"net/http"
)

func main() {

	db, err := database.Connection()

	if err != nil {
		log.Fatal(err)
	}

	studentRepository := repository.NewStudentRepository(db)

	studentService := service.NewStudentService(studentRepository)

	studentHandler := handler.NewStudentHandler(studentService)

	router := handler.SetupRoutes(studentHandler)

	http.ListenAndServe(":8080", router)

}