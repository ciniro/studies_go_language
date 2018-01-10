package main

import "fmt"

func main() {

//lendo valores
  fmt.Println("deseja continuar? 0=s/1=n")
  var output float64
  fmt.Scanf("%f", &output)
  fmt.Println(output)

//uso do IF
  if output == 0 {

//uso do FOR - separado com ponto virgula
    for i := 1; i <= 3; i++ {

      if i % 2 == 0 {
        fmt.Println("i é par")
      } else {
        fmt.Println("i é Impar")
      }

//uso do SWITCH
      switch i {
        case 1: {
          fmt.Println("Primeira instrução")
          fmt.Println("Segunda instrução")
        }
        case 2: fmt.Println("Instrução única")
        default: fmt.Println("Instrução padrão")
      }

      for j := 1; j <= 3; j++ {

        if j % 2 == 0 {
          fmt.Println("j é Divisivel por 2")
        } else if j % 3 == 0 {
          fmt.Println("j é Divisivel por 3")
        } else {
          fmt.Println("j Não é divisível por 2 ou 3")
        }

        fmt.Println(i, " - ", j)
      }
    }

  } else {
    fmt.Println("Vc encerrou")
  }


}
