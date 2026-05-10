package service

import (
	"api/internal/model"
	"api/internal/repository"
	"api/internal/utils"
	"strings"
)

func GetStudent() []model.Student {

	return repository.Find()

}

func PostStudent(students *model.StudentRequest) bool{

	for _, student := range students.Names {

		var newStudent model.Student
		
		newStudent.Name = student
		newStudent.Age = utils.RandAge()
		newStudent.Grades = utils.RandGrades(3)
		newStudent.Average = utils.CalculateAverage(newStudent.Grades)
		newStudent.Situation = utils.VerifySituation(newStudent.Average)

		repository.Create(newStudent)
	}

	return true
}

func DeleteStudent(studentName string) {

	for i, student := range repository.Students {

		if strings.EqualFold(student.Name, studentName) {

			repository.Students = append(repository.Students[:i], repository.Students[i + 1:]...)
			break

		}

	}

}