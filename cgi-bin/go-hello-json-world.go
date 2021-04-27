package main

import (
  "fmt"
  "net/http/cgi"
  "net/http"
  "time"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {
  header := w.Header()
  header.Set("Cache-Control", "no-cache")
  header.Set("Content-Type", "application/json")
  w.WriteHeader(200)

  fmt.Printf("{\n\t\"message\": \"Hello World From William and Chris From Golang\", \n")
  fmt.Printf("\t\"date\": \"%s\",\n", time.Now().String())
  fmt.Printf("\t\"currentIP\": \"%s\"\n}\n", r.RemoteAddr)
}

func main() {
  if err := cgi.Serve(http.HandlerFunc(RequestHandler)); err != nil {
    fmt.Println(err)
  }
}
