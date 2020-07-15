package main

import (
	"fmt"
	"log"
	"net/http"
)
import "github.com/gorilla/mux"

func main() {
	fmt.Println("Starting Up")
	r := mux.NewRouter()
	r.HandleFunc("/", TestRoute)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func TestRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello WORLD!\n"))
}