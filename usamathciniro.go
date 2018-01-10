package main

import "fmt"
import mc "mathciniro" //mc é um alias para mathciniro

func main() {
  xs := []float64{1,2,3,4}
  avg := mc.Average(xs)
  fmt.Println(avg)
}
