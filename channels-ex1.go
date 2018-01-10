package main

import ("fmt";"time";"strconv")

//chan é uma palavra reservada para definir um canal que deve ser procedido
//do tipo que esse canal transmite (string no caso)
//o operador de seta <- é usada para enviar e transmitir informacoes pelo
//canal
func ping(c chan string) {
  for i := 0; ; i++ {
    c <- "PING-" + strconv.Itoa(i) //envie a palavra PING para o canal
  }
}

func pong(c chan string) {
  for i := 0; ; i++ {
    c <- "PONG-" + strconv.Itoa(i) //envie a palavra PONG para o canal
  }
}

func printer(c chan string) {
  for {
    msg := <- c //receba do canal
    fmt.Println(msg) // poderia ser fmt.Println(<-c)
    time.Sleep(time.Second * 1)
  }
}

func main() {
  var c chan string = make(chan string)
  //duas goroutines vinculadas por um canal se sincronizam
  //pinger espera printer estar pronta para receber a mensagem e
  //so quando isso ocorrer a envia. No caso printer estara pronta para
  //receber a mensagem quando sleep terminar.
  go ping(c)
  go pong(c)
  go printer(c)

  var input string
  fmt.Scanln(&input)
}
