package main

import "fmt"

func main() {
  //criacao de um mapa de inteiros com indice literal
  x :=  make(map[string]int)

  x["idade"] = 33
  x["peso"] = 78

  fmt.Println(x)
  fmt.Println(x["idade"])
  fmt.Println(x["peso"])

//criacao de um mapa de strings com indice numerico inteiro
  y := make(map[int]string)
  y[0] = "ciniro"
  y[1] = "luis"
  y[3] = "joana"
  y[10] = "edmundo"

  fmt.Println(y)
  fmt.Println(y[3])
  //um mapa devolve vazio (string) ou zero (numero) quando o indice nao foi criado
  fmt.Println(y[2])

  delete(y, 3)

  fmt.Println(y)

  //pega o conteudo de um indice e retorna true/false se ele existe ou nao
  conteudo, existe := y[0]
 fmt.Println(conteudo, " - ", existe)

  if conteudo, existe := y[0]; existe {
    fmt.Println(conteudo, existe)
  }

  //criacao sintetica de mapas dentro de maps
  elementos := map[string]map[string]string {
      "H"   : map[string]string {
        "nome"   : "Hidrogenio",
        "estado" : "Gasoso",
      },

      "Cu"  : map[string]string {
        "nome"   : "Cobre",
        "estado" : "Solido",
      },
  }

  if conteudo, existe := elementos["Cu"]; existe {
    fmt.Println(conteudo["nome"], " - ", conteudo["estado"])
  }

  //encontrando o menor numero
  lista := []int{
    48,96,86,68,57,82,63,70,37,34,83,27,19,97,9,17,1,
  }

  menor := lista[0]
  for i := 0; i < len(lista); i++ {
    if lista[i] < menor {
      menor = lista[i]
    }
  }

  fmt.Println(menor)


}
