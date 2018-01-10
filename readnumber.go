package main

import "fmt"

func main() {
  fmt.Println("Entre com a temperatura em fareheint")
  var fareheint float64
  fmt.Scanf("%f", &fareheint)
  celsius := (fareheint - 32) * 5/9
  fmt.Println(celsius)
}
