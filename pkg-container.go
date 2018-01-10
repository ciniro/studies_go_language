package main

import ("fmt";"container/list")

func main() {
  //utilizando listas
  //criando uma lista duplamente ligada
  var x list.List
  x.PushBack(1) //adicionando um item na lista (ja inclui a ligacao de no)
  x.PushBack(2)
  x.PushBack(3)

  //do primeiro item ate que seja nulo o ultimo item va indo de item em item
  for e := x.Front(); e != nil; e = e.Next() {
    fmt.Println(e.Value.(int))
  }
}
