package main

import ("fmt"; "math")

//define um conjunto de metodos que um tipo devera ter se implementar essa
//interface
type Forma interface {
  area() float64
}

//criacao de uma forma que tem metodo area
//nao e necessario nenhuma palavra reservada como implements ou extends
//se um struct ja foi vinculado a um metodo que tenha a mesma assinatura de
//um metodo previsto na interface, entao ele implementa aquela interface
type Circulo struct {
  raio float64
}

//adicionando um metodo a um struct como um metodo de classe
func (c *Circulo) area() float64 {
  return math.Pi * (c.raio * c.raio)
}

type Triangulo struct {
  base, altura float64
}

//adicionando um metodo a um struct como um metodo de classe
func (t *Triangulo) area() float64 {
  return (t.base * t.altura)/2
}

//funcao para calcular a area de varias formas ao mesmo tempo sendo que neste
//caso as formas devem implementar a interface Forma
func totalAreas(f ...Forma) (total float64) {
  for _, valor := range f {
    total += valor.area()
  }

  return
}

//implementado agora um "MultiShape", uma colecao de formas
//ele cria uma variavel formas que contera um vetor com qualquer struct
//que implemente a interface Forma
type MultiShape struct {
  formas []Forma
}

//É possível implementar a interface inclusive diretamente do MultiShape
func (m *MultiShape) area() float64 {
  area := 0.0

  for _, valor := range m.formas {
    area += valor.area()
  }

  return area
}

func main() {
  c := Circulo{10}
  t := Triangulo{10, 30}
  //pode se chamar o metodo total areas que com a interface conhecendo
  //o metodo area simplesmente faz o calculo sem conhecer o tipo da forma
  fmt.Println(totalAreas(&c, &t))

  //criando um multishape também e possivel calcular a area adicionando
  //o proprio metodo de calcula das areas de todas as formas ao struct
  m := MultiShape {
    formas: []Forma {
      &c,
      &t,
    },
  }

  fmt.Println(m.area())

}
