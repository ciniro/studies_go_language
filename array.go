package main

import "fmt"

func main() {

  //declaracao de vetor
  var x [5]int
  x[4] = 100
  fmt.Println(x)

//vetores comecam em zero
  var m [5]float64
  m[0] = 10
  m[1] = 20
  m[2] = 30
  m[3] = 40
  m[4] = 50

//calculando uma media baseada em vetor
  var total float64 = 0

  for i := 0; i <= 4; i++ {
    total += m[i]
  }

//capturando o tamanho do vetor e convertendo int para float64
  fmt.Println(total/ float64(len(m)) )

//inclusao direta de elementos em um vetor
//se usar muitas linhas a ultima deve ser finalizada com virgula
  //m2 :=  [3]float64 {1,2,3}
  m2 :=  [3]float64 {
    1,
    2,
    3,}

  total2 := 0.0

//tipo de foreach, valor recebe cada elemento. O i é dispensado com o uso de
//um underscore
//este FOR é separado com virgula simples
  for _, valor := range m2 {
    total2 += valor
  }

  fmt.Println(total2/ float64(len(m2)) )
}
