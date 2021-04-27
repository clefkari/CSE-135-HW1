package main

import (
  "net/http"
  "net/http/cgi"
  "fmt"
)


func setHeader(w http.ResponseWriter, r *http.Request) string {
  header := w.Header()
  header.Set("Cache-Control", "no-cache")
  header.Set("Content-Type", "text/html")

  name := ""

  switch (r.Method) {
    case "POST":
      r.ParseForm()
      name = r.PostForm.Get("username")
      if name == "" { name = "No name given" }
      header.Set("Set-Cookie", "username="+name)
    default:
      cookie, _ := r.Cookie("username")
      name = cookie.Value
  }
  w.WriteHeader(200)
  return name
}

func addLinks() {
  fmt.Printf("<a href=/cgi-bin/go-sessions-2.cgi>Sessions Page 2</a><br/>")
  fmt.Printf("<a href=/go-cgiform.html>Golang CGI Form</a><br/>")
  fmt.Printf("<br/>")
}

func addDestroyButton() {
  fmt.Printf("<form action=\"/cgi-bin/go-destroy-session.cgi\" method=\"get\">")
  fmt.Printf("<button type=\"submit\">Destroy Session</button>")
  fmt.Printf("</form>")
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
  name := setHeader(w, r)

  fmt.Printf("<html><head><title>Golang Sessions</title></head><body>\n")
  fmt.Printf("<h1>Go Sessions Page 1</h1><br/>\n")
  fmt.Printf("<b>Name:</b> " + name + "<br/>")
  addLinks()
  addDestroyButton()
  fmt.Printf("</body></html>\n")
}


func main() {
  if err := cgi.Serve(http.HandlerFunc(requestHandler)); err != nil {
    fmt.Println(err)
  }
}


