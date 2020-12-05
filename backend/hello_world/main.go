package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Fprintln(os.Stdout, "Listening...")
	http.ListenAndServe(":8000", nil)
}
