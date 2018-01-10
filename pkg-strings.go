package main

import ("fmt";"strings")

func main() {
//encontrar uma string especifica em uma outra
fmt.Println(strings.Contains("Ciniro Nametala", "Name"))

//contar a quantidade de ocorrencias de uma string em outra
fmt.Println(strings.Count("Ciniro Nametala", "i"))

//confirmar se uma string se inicia com outra especifica
fmt.Println(strings.HasPrefix("Ciniro Nametala", "C"))

//confirmar se uma string termina com outra espeficica
fmt.Println(strings.HasSuffix("Ciniro Nametala", "a"))

//localizar a posicao de uma string em outra
//-1 se não encontrar a string
fmt.Println(strings.Index("Ciniro Nametala", "ni"))

//para concatenar strings separadas por uma separadora
fmt.Println(strings.Join([]string{"A","B","C"}, "-"))

//repetir uma string
fmt.Println(strings.Repeat("K", 10))

//substituir strings especificas com uma quantidade indicada
//-1 indica todas as substituicoes
fmt.Println(strings.Replace("Cxnxro fox ensxnar Go", "x", "i", -1))

//separar strings utilizando se uma string de separacao como ponto virgula
//gera um vetor de strings
fmt.Println(strings.Split("1;2;3;4;5", ";"))

//conversao entre caixas alta e baixa
fmt.Println(strings.ToLower("CINIRO"))
fmt.Println(strings.ToUpper("ciniro"))

}
