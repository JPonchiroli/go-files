package repository

import "api/internal/model"

func Find() model.Aluno {

	grades := []float64{1.0, 2.0, 3.0}

	return model.Aluno{Nome:"joao", Idade:20, Notas:grades, Media:4.0, Situacao:"test"}

}