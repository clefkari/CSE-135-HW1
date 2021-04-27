package main

import (
  "net/http"
  "net/http/cgi"
  "fmt"
)


func setHeader(w http.ResponseWriter) {
  header := w.Header()
  header.Set("Cache-Control", "no-cache")
  header.Set("Set-Cookie", "username=destroyed")
  header.Set("Content-Type", "text/html")
  w.WriteHeader(200)
}

func addLinks() {
  fmt.Printf("<a href=/cgi-bin/go-sessions-1.cgi>Sessions Page 1</a><br/>")
  fmt.Printf("<a href=/cgi-bin/go-sessions-2.cgi>Sessions Page 2</a><br/>")
  fmt.Printf("<a href=/go-cgiform.html>Golang CGI Form</a><br/>")
  fmt.Printf("<br/>")
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
  setHeader(w)

  fmt.Printf("<html><head><title>Go Session Destroyed</title></head><body>\n")
  fmt.Printf("<h1>Session Destroyed</h1><br/>\n")
  addLinks()
  fmt.Printf("</body></html>\n")
}


func main() {
  if err := cgi.Serve(http.HandlerFunc(requestHandler)); err != nil {
    fmt.Println(err)
  }
}



