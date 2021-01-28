package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", nil)
}

func helloWorld() {
	fmt.Println("Hello, world.")
}
