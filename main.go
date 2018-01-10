package main

import "fmt"

//para declarar varias variaveis escrevendo var apenas uma vez
var (
	nome string = "ciniro"
	idade = 33
	email string = "ciniro@gmail.com"
)

func main() {
	var x string = "Hello, World"
	fmt.Println(x)

//formas de declarar variaveis
//explicita
	var num0 int = 1
//implicita 1
	var num1 = 7
//implicita 2
	num2 := 2
//para sobreescrever um valor declarado com : não se deve usar : novamente
//pois seria o mesmo que estivesse declarando a mesma variavel duas vezes
	num2 = 3
	
	fmt.Println(num0 + num1 + num2)

	mostraNome()
}

//exemplo de funcao
func mostraNome() {
	fmt.Println(nome)
	fmt.Println(idade, " " + email)
}
