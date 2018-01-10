package main

import ("fmt";"time")
//esse programa exibe uma mensagem a cada um ou dois segundos passados
func main() {
  c1 := make(chan string)
  c2 := make(chan string)

  go func() {
    for {
      c1 <- "Passou mais UM segundo"
      time.Sleep(time.Second * 1)
    }
  }()

  go func() {
    for {
      c2 <- "Passaram-se mais DOIS segundos"
      time.Sleep(time.Second * 2)
    }
  }()


//select seleciona o canal que estiver pronto e executa a mensagem
//caso mais de um esteja pronto ele seleciona qualquer um aleatoriamente
//se nenhum estiver pronto ele espera ate que um esteja
  go func() {
    for {
      select {
        case msg1 := <- c1:
          fmt.Println(msg1)
        case msg2 := <- c2:
          fmt.Println(msg2)
        case <- time.After(time.Second):
          fmt.Println("Timeout")
        //default:
          //fmt.Print(".")
      }
    }
  }()

  var input string
  fmt.Scanln(&input)
}
