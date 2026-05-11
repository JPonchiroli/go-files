package utils

import (
	"math"
	"math/rand"
)

func RandAge() int {
  min, max := 14, 20
  age := min + rand.Intn(max - min)
  return age
}

func RandGrades(gradeQuantity int) []float64 {
  
  i := 0
  min, max := 0.0, 10.0
  grades := make([]float64,0)
  
  for i != gradeQuantity{
    grade := min + rand.Float64() * (max - min)
    grade = math.Round(grade*100) / 100
    grades = append(grades, grade)
    i++
  }
  
  return grades
  
  
}

func CalculateAverage(grades []float64) float64 {
  
  sum, average := 0.0, 0.0
  
  
  for _, nota := range grades {
    sum += nota
  }
  
  average = sum / float64(len(grades))
  average = math.Round(average*100) / 100
  
  return average
  
}

func VerifySituation(average float64) string {
  
  situation := ""
  
  
  if (average <= 3.0) {
    
    situation = "Reproved"
    
  } else if (average <= 6.0) {
    
    situation = "In exam"
    
  } else {
    
    situation = "Approved"
    
  }
  
  return situation
  
}