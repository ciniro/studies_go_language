package main

import ("fmt"; "flag"; "math/rand")
//flags sao argumentos que podem ser passados ao programa no momento de
//sua chamada via linha de comando no cmd. As flags sao capturadas pelo
//codigo e utilizadas. Nesse exemplo a flag servira para definir o maximo
//valor para gerar um numero aleatorio entre 0 e max
//chamada: go run pkg-flag.go -max=100
func main() {
  //criacao da flag
  maxp := flag.Int("max", 6, "o valor maximo")
  //converte a mesma
  flag.Parse()
  //gera o numero aleatorio
  fmt.Println(rand.Intn(*maxp))
}
