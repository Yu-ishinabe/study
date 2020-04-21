package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<!DOCTYPE html>")
	fmt.Fprintln(w, "<html lang='ja'>")
	fmt.Fprintln(w, "<head>")
	fmt.Fprintln(w, "<meta charset='UTF-8'>")
	fmt.Fprintln(w, "<title>TITLE</title>")
	fmt.Fprintln(w, "</head>")
	fmt.Fprintln(w, "<body>")
	fmt.Fprintln(w, "<h1>hello GO!</h1>")
	fmt.Fprintln(w, "</body>")
	fmt.Fprintln(w, "</html>")
}

func main() {
	http.HandleFunc("/", home)
	if err := http.ListenAndServe(":8080", nil): err != nil {
		log.Fatal("ListenAndServe failed.", err)
	}
}
