package service

import (
	"api/internal/model"
	"api/internal/repository"
	"api/internal/utils"
)

func GetStudent() []model.Student {

	return repository.Find()

}

func PostStudent(student *model.Student) model.Student{

	student.Age = utils.RandAge()
	student.Grades = utils.RandGrades(3)
	student.Average = utils.CalculateAverage(student.Grades)
	student.Situation = utils.VerifySituation(student.Average)

	repository.Create(*student)

	return *student
}