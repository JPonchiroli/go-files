package service

import (
	"api/internal/model"
	"api/internal/repository"
)

func GetStudent() model.Aluno {

	return repository.Find()

}