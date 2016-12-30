package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"fmt"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	log.Print("Starting webserver...")
	log.Fatal(http.ListenAndServe("localhost:8090", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
	w.WriteHeader(http.StatusOK)
}