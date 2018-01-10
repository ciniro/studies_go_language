package main

import ("fmt"; "os"; "io/ioutil"; "path/filepath")

func main() {
  //abrindo e lendo um arquivo texto
  bs, err := ioutil.ReadFile("test.txt")

  if err != nil {
    return
  }

  str := string(bs)
  fmt.Println(str)

  //criando um arquivo
  file, err := os.Create("test2.txt")

  if err != nil {
    return
  }

  defer file.Close()

  file.WriteString("Incluindo texto no arquivo 2")

  //obtendo o conteudo de um diretorio inteiro
  dir, err := os.Open("C:\\Users\\ciniro\\goprojects")

  if err != nil {
    return
  }

  defer dir.Close()

  fileInfos, err := dir.Readdir(-1)

  if err != nil {
    return
  }

  for _, valor := range fileInfos {
    fmt.Println(valor.Name())
  }

  //funcao walk para percorrer diretorios e subdiretorios
  filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
    fmt.Println(path)
    return nil
  })
}
