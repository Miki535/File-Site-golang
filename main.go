package main

import (
	"html/template"
	"net/http"
)

var tpl = template.Must(template.ParseFiles("templates/index.html"))

func main() {
	http.HandleFunc("/", Osn)
	_ = http.ListenAndServe(":8080", nil)
}

func Osn(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		filename := r.FormValue("filename")
		info := r.FormValue("info")

	}
	_ = tpl.Execute(w, nil)
}
