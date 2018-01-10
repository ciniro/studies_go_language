package main

import "fmt"

func main() {
  //forma 1 de criar uma slice com 10 posicoes onde so se usa 5
  x := make([]float64, 5, 10)
  x[0] = 3
  x[4] = 1

  fmt.Println(x)

  //adicionando novos elementos numa slice
  x = append(x, 22, 23, 24)

  fmt.Println(x)

  //forma 2 de criar uma slice com 5 posicoes onde se usa as 5 mas para
  //mostrar ele oferece o recurso :, como no R
  //isto e a criacao de um array
  arr := [5]float64 {1,2,3,4,5}

  //ja isso e a criacao de uma slice a partir do array
  z := arr[:]
  fmt.Println(z)

  z = arr[2:]
  fmt.Println(z)

  z = arr[2:4]
  fmt.Println(z)

  //nao vai funcionar pois e um array e nao um slice
  // m := append(arr, 6, 7)
  m := append(z, 6, 7)

  fmt.Println(m)

}
