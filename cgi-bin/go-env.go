package main

import (
  "net/http"
  "net/http/cgi"
  "fmt"
  "os"
)


func requestHandler(w http.ResponseWriter, r *http.Request) {
  header := w.Header()
  header.Set("Cache-Control", "no-cache")
  header.Set("Content-Type", "text/html")
  w.WriteHeader(200)

  fmt.Printf("<html><head><title>Environment Variables</title></head>\n")
  fmt.Printf("<body><h1 align=center>Environment Variables</h1><hr/>\n")

  for _, env_var := range os.Environ() {
    fmt.Printf(env_var + "<br/>")
  }

  fmt.Println("</body></html>")

}


func main() {
  if err := cgi.Serve(http.HandlerFunc(requestHandler)); err != nil {
    fmt.Println(err)
  }
}
