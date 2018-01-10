package main

import "errors"

//criando um novo tipo de erro
func main() {
  err := errors.New("Mensagem relacionada ao novo erro")
}
