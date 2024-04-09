package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
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
		FILE(filename, info)

	}
	tpl.Execute(w, nil)
}

func FILE(FILEName string, INFO string) {
	file, err := os.Create(FILEName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(INFO)
	if err != nil {
		return
	}
}
