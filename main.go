package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
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
	userInfo, err := user.Current()
	if err != nil {
		fmt.Println(err)
		return
	}

	downloadsPath := filepath.Join(userInfo.HomeDir, "Downloads")

	filepath := filepath.Join(downloadsPath, FILEName)

	file, err := os.Create(filepath)
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
