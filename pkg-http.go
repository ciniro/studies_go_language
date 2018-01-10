package main

import ("net/http"; "io")

func hello(res http.ResponseWriter, req *http.Request) {
  res.Header().Set (
    "Content-Type",
    "text/html",
  )

  html := "<DOCTYPE html><html><head><title>Hello, World!</title></head><body>Hello, World!</body></html>"

  io.WriteString (res,html)
}

func main() {
  http.HandleFunc("/hello", hello)
  http.ListenAndServe(":9000", nil)
}
