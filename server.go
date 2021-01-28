package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})

	http.Handle("/", http.FileServer(http.Dir("")))

	http.ListenAndServe(":8080", nil)
}

func helloWorld() {
	fmt.Println("Hello, world.")
}
