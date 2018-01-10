package main

import ("fmt"; "strconv")

//criando dois struct simples
type Amigo struct {
  nome string
}

type Arma struct {
  tipo string
  qtdeBalas int
}

//criando um struct e adicionando nele o struct criado anteriormente, alem
//de novos campos
type Pessoa struct {
  Amigo
  Arma
  nome, email string
  idade int
}

//adicionando um metodo ao struct criado
func (p *Pessoa) calcular(x int, y int) string {
  return "Olá! Eu posso calcular, " + strconv.Itoa(x) + " mais " + strconv.Itoa(y) + " é igual a " + strconv.Itoa(x + y) + "."
}

//criando um terceiro struct adicionado dos dois anteriores por meio de um
type Androide struct {
  Pessoa
  modelo string
}

//adicionando mais um metodo ao ultimo e final struct que compoe todos
//aqui deixei a variavel de retorno propositadamente definida antes, sem motivo
func (a *Androide) falar() (fala string) {
  fala = "Olá! Eu posso falar! Vejam só: Meu nome é " + a.nome + ". Meu email é " + a.email + ". Minha idade é " + strconv.Itoa(a.idade) + ". Meu modelo é " + a.modelo + ". Isso não é legal!?"
  return
}


func main() {
  //o uso de structs dessa forma lembra classes de POO
  c3po := new(Androide)
  c3po.nome = "C3PO"
  c3po.email = "c3po@starwars.com"
  c3po.idade = 200
  c3po.modelo = "Legacy"
  c3po.Amigo.nome = "R2D2"
  c3po.Arma.tipo = "Laser"
  c3po.Arma.qtdeBalas = 25

  fala := c3po.falar()
  calculo := c3po.calcular(2,3)

  fmt.Println(fala)
  fmt.Println(calculo)
  fmt.Println(c3po.Amigo.nome)
  fmt.Println(c3po.Arma)
}
