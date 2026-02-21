package main

import (
	"log"
	"net/http"
	"fmt"
)

func home (w http.ResponseWriter, r *http.Request ){
	w.Write([]byte("Hello from snippetbox123\n"))
}

func josh (w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello world, %s!", r.URL.Path[1:])
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/josh", josh)

	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}