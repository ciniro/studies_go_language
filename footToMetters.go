package main

import "fmt"
//declaracao de constantes
const pe float64 = 0.3042

func main() {
  fmt.Println("Entre com a qtde de pes")
  var pes float64
  fmt.Scanf("%f", &pes)
  metros := pes * pe
  fmt.Println(metros)
}
