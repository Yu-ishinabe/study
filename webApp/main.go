package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type Greeting struct {
	Message string
	Time    time.Time
}

func home(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/sample.html")
	m := Greeting{
		Message: "Hello Go lang!",
		Time:    time.Now(),
	}
	t.Execute(w, m)
}

func main() {
	// cssの読込
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css/"))))
	http.HandleFunc("/sample", home)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe failed.", err)
	}
}
