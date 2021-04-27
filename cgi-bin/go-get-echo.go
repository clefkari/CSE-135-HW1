package main

import (
  "net/http"
  "net/http/cgi"
  "fmt"
  "encoding/json"
  "net/url"
  "strings"
)


func setHeader(w http.ResponseWriter) {
  header := w.Header()
  header.Set("Cache-Control", "no-cache")
  header.Set("Content-Type", "text/html")
  w.WriteHeader(200)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
  setHeader(w)

  fmt.Printf("<html><head><title>GET query string</title></head><body>\n")
  fmt.Printf("<h1 align=center>GET query string</h1>\n")
  fmt.Printf("<hr/>")
  fmt.Printf("Raw query string: " + r.URL.Query().Encode() + "<br/>")

  query_string, _ := url.ParseQuery(r.URL.Query().Encode())
  js, _ := json.Marshal(query_string)
  fmt.Printf("Formatted query string: " + strings.ReplaceAll(string(js), ",", ", "))

  fmt.Printf("</body></html>\n")
}


func main() {
  if err := cgi.Serve(http.HandlerFunc(requestHandler)); err != nil {
    fmt.Println(err)
  }
}
