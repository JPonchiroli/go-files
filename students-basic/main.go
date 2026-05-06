package main
import "fmt"
import "math"
import "math/rand/v2"


type Aluno struct {
  nome string
  idade int
  notas []float64
  media float64
  situacao string
}

func main() {
  
  alunos := populaArrayAlunos()
  
  for _, aluno := range alunos {
    exibirBoletim(aluno)
  }
  
}

func populaArrayAlunos() []Aluno {
  
  nomes := []string{"Joao", "Maria", "Pedro", "Ana", "Lucas", "Carla", "Rafael", "Juliana"}
  alunos := make([]Aluno,0)
  
  for _, nome := range nomes {
    idade := sorteiaIdade()
    notas := sorteiaNotas(3)
    media := calculaMedia(notas)
    situacao := verificaSituacao(media)
    alunos = append(alunos, Aluno{nome, idade, notas, media, situacao})
  }
  
  return alunos
  
}

func sorteiaIdade() int {
  min, max := 14, 20
  idade := min + rand.IntN(max - min)
  return idade
}

func sorteiaNotas(quantidadeNotas int) []float64 {
  
  i := 0
  min, max := 0.0, 10.0
  notas := make([]float64,0)
  
  for i != quantidadeNotas{
    nota := min + rand.Float64() * (max - min)
    nota = math.Round(nota*100) / 100
    notas = append(notas, nota)
    i++
  }
  
  return notas
  
  
}

func calculaMedia(notas []float64) float64 {
  
  soma, media := 0.0, 0.0
  
  
  for _, nota := range notas {
    soma += nota
  }
  
  media = soma / float64(len(notas))
  media = math.Round(media*100) / 100
  
  return media
  
}

func verificaSituacao(media float64) string {
  
  situacao := ""
  
  
  if (media <= 3.0) {
    
    situacao = "Reprovado"
    
  } else if (media <= 6.0) {
    
    situacao = "Em exame"
    
  } else {
    
    situacao = "Aprovado"
    
  }
  
  return situacao
  
}

func exibirBoletim(aluno Aluno) {
  
  infoAluno := fmt.Sprintf("Aluno: %s   Idade: %d", aluno.nome, aluno.idade) 
  situacao := fmt.Sprintf("Média: %v   Situação: %s", aluno.media, aluno.situacao) 
  notas := fmt.Sprintf("Notas: %v", aluno.notas)
  
  fmt.Println("----------------------------------")
  fmt.Println(infoAluno)
  fmt.Println(situacao)
  fmt.Println(notas)
  fmt.Println("----------------------------------")
}