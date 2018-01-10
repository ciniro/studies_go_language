package main

import ("fmt";"io/ioutil";"net/http";"time")

//funcao que recebe a URL e guarda no canal c o seu tamanho
func analisaURL(url string) int {
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
  return len(bs)
}

func preencheCanal(urls []string) [5]int {
  inicio := time.Now().UTC()
  i := 0
  var t [5]int

  for _, url := range urls {
    //sera criada uma goroutine para cada url armazenada
    //e todas seram examinadas simultaneamente
    t[i] = analisaURL(url)
    i++

  }

  time.Sleep(100 * time.Millisecond)
  decorrido := time.Since(inicio) - (100 * time.Millisecond)
  fmt.Println("Tempo decorrido: ", decorrido)

  return t
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
  var tamanhos [5]int

  tamanhos = preencheCanal(urls)

  //so continua a execucao aqui quando as rotinas do for estiverem
  //finalizadas
  maiorpagina := urls[0]
  maioratual := tamanhos[0]

  for i := 0; i < len(urls); i++ {
    if tamanhos[i] > maioratual  {
      maiorpagina = urls[i]
      maioratual = tamanhos[i]
    }
  }


  fmt.Println("A maior pagina é ", maiorpagina)
}
