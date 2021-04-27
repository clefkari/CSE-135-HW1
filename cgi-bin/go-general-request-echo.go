package main

import (
  "net/http"
  "net/http/cgi"
  "fmt"
  "strings"
  "io"
)


func setHeader(w http.ResponseWriter) {
  header := w.Header()
  header.Set("Cache-Control", "no-cache")
  header.Set("Content-Type", "text/html")
  w.WriteHeader(200)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
  setHeader(w)

  fmt.Printf("<html><head><title>General Request Echo</title></head><body>\n")
  fmt.Printf("<h1 align=center>General Request Echo</h1>\n")
  fmt.Printf("<hr/>")
  fmt.Printf("Protocol: " + r.Proto + "<br/><br/>")
  fmt.Printf("Method: " + r.Method + "<br/><br/>")
  fmt.Printf("Message Body: ")
  body := new(strings.Builder)
  io.Copy(body, r.Body)
  fmt.Printf(body.String())
  fmt.Printf("</body></html>\n")
}


func main() {
  if err := cgi.Serve(http.HandlerFunc(requestHandler)); err != nil {
    fmt.Println(err)
  }
}


