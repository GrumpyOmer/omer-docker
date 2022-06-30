package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("this is a golang service")
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8765", nil)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
