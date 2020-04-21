package main

impot (
	"net/http"
	"fmt"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "test message")
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}
