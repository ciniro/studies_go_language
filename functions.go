package main

import "fmt"
//para usar uma variavel de escopo geral a mesma pode ser declarada fora
//das funcoes, como isso:
var nome string = "Ciniro"

func main() {
  vetor := []float64{1,2,3,4,5}
  fmt.Println(average(vetor))
  fmt.Println(somaString("Nametala"))
  fmt.Println(retorno(10))

  x, y, z, s := retornomultiplo()
  fmt.Println(x)
  fmt.Println(y)
  fmt.Println(z)
  fmt.Println(s)

  vint, vstring := variadica("Ciniro", "Aparecido", "Leite", "Nametala")
  fmt.Println(vint)
  fmt.Println(vstring)

  //e possivel criar uma funcao dentro de outra, o que é chamado de Closure
  valor := 0
  incremento := func() int {
    valor++
    return valor
  }
  fmt.Println(incremento())
  fmt.Println(incremento())
  fmt.Println(incremento())

  //funcao gerando funcao com valor local dentro da funcao
  proximoPar := geraNumerosPares()
  fmt.Println(proximoPar())
  fmt.Println(proximoPar())
  fmt.Println(proximoPar())
  fmt.Println(proximoPar())
  fmt.Println(proximoPar())

  //usando recursao
  fmt.Println(fatorial(4))
}

//criando uma funcao
//calcula media
func average(xs []float64) float64 {
  total := 0.0

  for _, valor := range xs {
    total += valor
  }

  return total / float64(len(xs))
}

//concatena duas strings
func somaString(valor string) string {
  return nome + " " + valor
}

//o retorno pode estar explicitamente vinculado a uma variavel especifica
//dentro da funcao, neste caso a variavel ja esta declarada
func retorno(valor int) (resposta int) {
  resposta = 1 + valor
  return
}

//muito util. Go permite a devolucao de quantos valores de retorno forem
//necessarios
func retornomultiplo() (int, int, int, string) {
  return 10, 20, 30, "quarto valor"
}

//funcao variadica e uma funcao que com o uso de reticencias permite passar
//como parametro qualquer quantidade de valores
//o parametro tem que ser unico e de nome args
func variadica(args ...string) (int, string) {
  texto := ""
  for _, valor := range args {
    texto += " " + valor
  }

  soma := 10
  return soma, texto
}

//e possivel criar uma funcao que devolva uma funcao. nesta funcao pode-se
//tambem deixar variaveis locais que serao usadas como parametro para a funcao
//nova criada. Veja que o retorno desta funcao é uma outra funcao
func geraNumerosPares() func() uint {
  i := uint(0) //gera um uint de valor zero, podia ser outro valor qualquer
  return func() (ret uint) {
    ret = i
    i += 2
    return
  }
}

//recursao para calculo de fatorial
func fatorial(x uint) uint {
  if (x == 0) {
    return 1
  }

  return x * fatorial(x-1)
}
