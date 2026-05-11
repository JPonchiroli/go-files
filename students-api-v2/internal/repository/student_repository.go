package repository

import "api/internal/model"

var Students []model.Student

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