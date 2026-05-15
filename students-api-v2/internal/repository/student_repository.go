package repository

import (
	"api/internal/model"
	"database/sql"
)

type StudentRepository struct {
	DB *sql.DB
}

func NewStudentRepository(db *sql.DB) *StudentRepository {

	return &StudentRepository{
		DB: db,
	}

}

func Find() []model.Student {

	students := make([]model.Student,0)

	for _, studentsRegistered := range Students {
		students = append(students, studentsRegistered)
	}

	return students

}

func Create(student model.Student) {

	Students = append(Students, student)

}