package main

import (
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("sample.html")
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/sample", home)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe failed.", err)
	}
}
