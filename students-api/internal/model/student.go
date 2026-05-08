package model

type Student struct {
  Id        int       `json:"id"`     
  Name      string    `json:"name"`
  Age       int       `json:"age"`
  Grades    []float64 `json:"grades"`
  Average   float64   `json:"average"`
  Situation string    `json:"situation"`
}