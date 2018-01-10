package main

import ("fmt";"time";"math/rand")

func funcao(n int) {
  for i := 0; i < 10; i++ {
    fmt.Println(n, ":", i)
  }
}

func funcao2(n int) {
  for i := 1; i <= 2; i++ {
    fmt.Println(n, ":", i)

    tempo:= time.Duration(rand.Intn(250))
    time.Sleep(time.Millisecond * tempo)
  }
}

//goroutines sao alocadas a um processador para serem executadas em separado
//o programa entao continua a sua execucao sem esperar a funcao terminar
//por isso existe no exemplo abaixo o leitor de tela, sem ele nao daria
//tempo de ver o resultado da rotina go criada
func main() {
  fmt.Println("Digite o exemplo que quer ver:")
  var input int
  fmt.Scanln(&input)

  switch input {
    case 1: {
      //exemplo 1 - executando uma goroutine
      go funcao(0)
      fmt.Scanln(&input)
    }
    case 2: {
      //exemplo 2 - executando varias gouroutines
      for i := 1; i <= 10; i++ {
        go funcao(i)
      }
      fmt.Scanln(&input)
    }
    case 3: {
      //exemplo 2 - executando varias gouroutines com um tempo
      //randomico de delay em cada 
      for i := 1; i <= 5; i++ {
        go funcao2(i)
      }
      fmt.Scanln(&input)
    }
    default: fmt.Println("Exemplo invalido")
  }
}
