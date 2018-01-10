package main

import ("fmt";"io";"os";"hash/crc32")

//funcao que recebe um arquivo e devolve o hash em uint32 e um erro
//(nulo ou nao)
func geraHash(arquivo string) (uint32, error) {
  //abre o arquivo
  f, err := os.Open(arquivo)
  if err != nil {
    return 0, err
  }

  defer f.Close()

  //gera um hash vazio
  h := crc32.NewIEEE()
  //recebe o hash baseado no arquivo
  _, err = io.Copy(h, f) // o parametro _ guardaria os bytes se fossem usados

  if err != nil {
    return 0, err
  }

  return h.Sum32(), nil //seta o uint32 com o valor do hash e o erro com nil
}

func main() {
  f1, err := geraHash("test.txt")
  if err != nil {
    return
  }

  f2, err := geraHash("test2.txt")
  if err != nil {
    return
  }

  fmt.Println(f1, f2, f1 == f2)
}
