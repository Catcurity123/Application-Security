package main

import (
	"fmt"
	"net/http"
)

func hello (w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe("localhost:8080", nil)
}