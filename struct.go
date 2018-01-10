package main

import ("fmt"; "math")

//criacao de um struct, adicionando campos, como propriedades de classe
type Circulo struct {
  x, y, raio float64
}

//adicionando um metodo a um struct como um metodo de classe
func (c *Circulo) area() float64 {
  return math.Pi * (c.raio * c.raio)
}

//funcao que calcula a area do circulo recebendo um circulo como parametro
func areaCirculo(c Circulo) float64 {
  return math.Pi * (c.raio * c.raio)
}

func main() {
//inicializando um struct como referencia
 //var c Circulo
 //c := new(Circulo)
 //c := Circulo{x: 0, y: 0, raio: 5}
 //c := &Circulo{0, 0, 5}
 c := Circulo{0, 0, 5}

 //acessando os campos de um struct
 fmt.Println(c.x, c.y, c.raio)
 //passando um struct como parametro
 fmt.Println(areaCirculo(c))

 //chamando um metodo adicionado ao struct
 fmt.Println(c.area())
}
