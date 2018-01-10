//para criar um pacote é preciso primeiro criar um diretorio src dentro do
//diretorio indicado para GOPATH. Nesse diretorio src podese criar pastas
//respectivas para cada pacote a ser implementado. Nessas pastas devera ser
//incluido o arquivo do pacote, como esse abaixo.

//O NOME DO PACOTE TEM QUE SER O MESMO DA PASTA ONDE O MESMO SE ENCONTRA

//IMPORTANTE!
//a visibilidade da funcao é determinada assim:
//a) Se começa com letra maiuscula ("Average") entao sera publica
//b) Se começa com letra minuscula ("average") entao sera privada

//pode colocar um comentario de documentacao se este estiver
//imediatamente anterior ao nome do pacote ou funcao
//para ver a documentacao tem duas formas:
//a) Console: godoc [NOMEDOPACOTE]
//b) Web: godoc -http=":6060" e acessando a seguinte url depois
//http://localhost:6060/pkg/

//pacote matematico criado pelo Ciniro
package mathciniro

//funcao que faz calculo de medias baseado em um vetor de floats
func Average(xs []float64) float64 {
  total := float64(0)
  for _, x := range xs {
    total += x
  }

  return total/float64(len(xs))
}

//apos criado o arquivo, devese acessar via terminal a pasta onde o arquivo
//esta, na sequencia usase o comando go install [NOMEDOPACOTE]. Sera entao
//criada automaticamente uma pasta chamada pkg e dentro dela outra chamada
//windows_amd64. Nesta estara o codigo-objeto compilado.

//Para usar o pacote basta importar da pasta GOPATH como em import "mathciniro"
