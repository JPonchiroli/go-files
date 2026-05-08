package service

import (
	"api/internal/model"
	"api/internal/repository"
)

func GetStudent() model.Student {

	return repository.Find()

}