//para execucao de testes devese criar um arquivo de mesmo nome do pacote
//que se deseja testar, contudo, o arquivo deve ter nome terminado com _test.go.
//Go nao executa e nem compila essa arquivo pois ja sabe por causa do _test que
//este é um arquivo exclusivamente para testar um pacote.

//O nome do pacote deve ser o mesmo no _test.
package mathciniro

import "testing"

type teste struct {
  valores []float64
  media float64
}

var conjuntodetestes = []teste{
  {[]float64{1,2}, 1.5},
  {[]float64{1,1,1,1,1,1,1,1},1},
  {[]float64{-1,1},0},
}

//A funcao a ser testada deve ser iniciada com Test[NOMEDAFUNCAO]
func TestAverage(t *testing.T) {
  for _, testecorrente := range conjuntodetestes {
    resultado := Average(testecorrente.valores)

    if resultado != testecorrente.media {
      t.Error (
        "Para", testecorrente.valores,
        "espera-se", testecorrente.media,
        "obteve-se", resultado,
      )
    }
  }
}
