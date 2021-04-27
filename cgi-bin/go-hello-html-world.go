package main

import (
  "net/http"
  "net/http/cgi"
  "time"
  "fmt"
)


func requestHandler(w http.ResponseWriter, r *http.Request) {
  header := w.Header()
  header.Set("Cache-Control", "no-cache")
  header.Set("Content-Type", "text/html")
  w.WriteHeader(200)

  fmt.Printf("<html><head><title>Hello Golang CGI</title></head><body>\n")
  fmt.Printf("<h1 align=center>Hello HTML Golang World</h1>\n")
  fmt.Printf("<hr/>")
  fmt.Printf("Hello World From William and Chris <br/>")
  fmt.Printf("This program was generated at: %s <br/>", time.Now().Format("01-02-2006 15:04:05"))
  fmt.Printf("Your current IP address is: %s <br/>", r.RemoteAddr)
  fmt.Printf("</body></html>\n")
}


func main() {
  if err := cgi.Serve(http.HandlerFunc(requestHandler)); err != nil {
    fmt.Println(err)
  }
}
