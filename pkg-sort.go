package main

import ("fmt"; "sort")

type Pessoa struct {
  nome string
  idade int
}

type PorNome []Pessoa

func (ps PorNome) Len() int {
  return len(ps)
}

func (ps PorNome) Less(i, j int) bool {
  return ps[i].nome < ps[j].nome
}

func (ps PorNome) Swap(i, j int) {
  ps[i], ps[j] = ps[j], ps[i]
}

func main() {
  criancas := []Pessoa {
    {"Joao", 9},
    {"Maria", 10},
    {"Ana Julia", 17},
  }

  sort.Sort(PorNome(criancas))
  fmt.Println(criancas)

  n1 := "Ciniro"
  n2 := "Aparecido"

  fmt.Println(n1 > n2)
}
