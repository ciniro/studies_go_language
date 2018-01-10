package main

import ("fmt";"io/ioutil";"net/http";"time")

//define o tipo que armazenara a url e o seu tamanho
type Pagina struct {
  URL string
  Size int
}

//funcao que recebe a URL e guarda no canal c o seu tamanho
func analisaURL(url string, c chan Pagina) {
  //faz a requisicao a pagina
  res, err := http.Get(url)

  if err != nil {
    panic(err)
  }

  defer res.Body.Close()
  //calcula o tamanho do corpo da pagina
  bs, err := ioutil.ReadAll(res.Body)

  if err != nil {
    panic(err)
  }

  //quando o calculo do tamanho estiver terminado ele é entao
  //enviado para o canal, assim que finalizar
  //perceba que o canal é uma especie de hub que armazena os resultados
  //das goroutines assim que estas terminam de ser executadas
  c <- Pagina{
    URL: url,
    Size: len(bs),
  }
}

func preencheCanal(urls []string, c chan Pagina) {
  inicio := time.Now().UTC()

  for _, url := range urls {
    //sera criada uma goroutine para cada url armazenada
    //e todas seram examinadas simultaneamente
    go analisaURL(url, c)
  }

  time.Sleep(100 * time.Millisecond)
  decorrido := time.Since(inicio) - (100 * time.Millisecond)
  fmt.Println("Tempo decorrido: ", decorrido)
}

func main() {
  //cria uma lista das urls a serem examinadas
  urls := []string{
    "http://www.apple.com",
    "http://www.amazon.com",
    "http://www.google.com",
    "http://www.microsoft.com",
    "http://www.ciniro.net",
  }

  //cria um canal que so armazenara o tipo homepagesize
  canal := make(chan Pagina)

  preencheCanal(urls, canal)

  //so continua a execucao aqui quando as rotinas do for estiverem
  //finalizadas
  var maiorpagina Pagina

  for range urls {
    resultado := <- canal

    if resultado.Size > maiorpagina.Size {
      maiorpagina = resultado
    }
  }

  fmt.Println("A maior pagina é ", maiorpagina.URL)
}
