package main

import (
  "net/http"
  "net/http/cgi"
  "fmt"
)


func setHeader(w http.ResponseWriter) {
  header := w.Header()
  header.Set("Cache-Control", "no-cache")
  header.Set("Content-Type", "text/html")
  w.WriteHeader(200)
}

func addLinks() {
  fmt.Printf("<a href=/cgi-bin/go-sessions-1.cgi>Sessions Page 1</a><br/>")
  fmt.Printf("<a href=/go-cgiform.html>Golang CGI Form</a><br/>")
  fmt.Printf("<br/>")
}

func addDestroyButton() {
  fmt.Printf("<form action=\"/cgi-bin/go-destroy-session.cgi\" method=\"get\">")
  fmt.Printf("<button type=\"submit\">Destroy Session</button>")
  fmt.Printf("</form>")
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
  setHeader(w)

  fmt.Printf("<html><head><title>Golang Sessions</title></head><body>\n")
  fmt.Printf("<h1>Go Sessions Page 2</h1><br/>\n")
  cookie, _ := r.Cookie("username")
  fmt.Printf("<b>Name:</b> " + cookie.Value + "<br/>")
  addLinks()
  addDestroyButton()
  fmt.Printf("</body></html>\n")
}


func main() {
  if err := cgi.Serve(http.HandlerFunc(requestHandler)); err != nil {
    fmt.Println(err)
  }
}


