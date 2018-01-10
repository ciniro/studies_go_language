package main

import "fmt"

func main() {
//PONTEIROS
//E possivel usar ponteiros para passar variaveis nao por referencia mas por
//valor para uma funcao qualquer. Um ponteiro declara explicitamente que voce
//quer acessar diretamente o local na memoria onde aquela varivel esta
//armazenada, nesse caso, voce estara inflingindo o escopo
x := 50
alteraComPonteiro(&x)
fmt.Println(x)

//é possivel também criar apenas uma referencia como ponteiros (endereco)
z := new(int)
alteraComPonteiro(z)
fmt.Println(z) //printando o endereco
fmt.Println(*z) //printando o conteudo do endereco

//DEFER
  //diferimento de funcao: Em go e possivel obrigar a chamada de uma funcao
  //no final de uma funcao qualquer, para fechar uma conexao com banco ou um
  //arquivo por exemplo. A funcao diferida sera executada antes do retorno e
  //mesmo se algum "panic" acontecer. Ela nao obedece a sequencia de chamada.
  defer textofinal()
  texto2()
  texto1()

//PANIC RECOVER
//chamar a funcao panic interrompe a execucao do codigo naquele ponto e mostra
//tambem uma serie de informacoes sobre o ponto de parada. Da pra usar panic
//para exibir conteudo de variaveis, como die no php. Para nao mostrar as
//informacoes e desviar o codigo para uma outra funcao, deve-se usar RECOVER
//junto com DEVER. RECOVER é uma funcao que captura o valor exibido no panic.
//Esse valor pode ser retornado para uma string e mostrado na tela. Se apenas
//chamar recover() entao o conteudo de panic não é mostrado, mas a funcao que
//contem o mesmo, chamada com defer executara do mesmo jeito.

  defer recuperacao()
  panic("panico")
  texto1() //devido ao panico este comando nao sera executado
}

//funcoes para exemplo de diferimento
func texto1() {
  fmt.Println("texto 1")
}

func texto2() {
  fmt.Println("texto 2")
}

func textofinal() {
  fmt.Println("texto final")
}

//funcao de recuperacao para entendimento de panic RECOVER
func recuperacao() {
  //recover()
  str := recover()
  fmt.Println(str)
}

//funcao que altera o conteudo de uma variavel qualquer inteira com ponteiro
//forma de passar variavel por valor e nao por referencia
func alteraComPonteiro(ponteiro *int) {
  *ponteiro = 25
}
