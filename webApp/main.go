package main

import (
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/sample.html")
	t.Execute(w, nil)
}

func main() {
	// cssの読込
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css/"))))
	http.HandleFunc("/sample", home)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe failed.", err)
	}
}
