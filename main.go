package main

import (
  	"net/http"
	"html/template"
	"log"
)

func main() {

  h1 := func (w http.ResponseWriter, r *http.Request){
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, nil)
  }
  http.HandleFunc("/", h1)

  log.Fatal(http.ListenAndServe(":8000", nil))
}