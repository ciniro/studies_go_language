package main

import ("fmt"; "crypto/sha1")

func main() {
//cria um hash vazio baseado em sha1
//esse metodo é igual ao usado no hash pois o metodo New implementa a mesma
//interface hash.Hash
  h := sha1.New()
  //converte senha para sha1 neste hash
  h.Write([]byte("senha"))
  bs := h.Sum([]byte{})
  //usa uma fatia pois 160 bits não caberiam em um hash que so aceita 32
  fmt.Println(bs)
}
