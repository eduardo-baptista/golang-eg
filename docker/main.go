package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Docker test is working")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handleRoot)

	log.Fatal(http.ListenAndServe(":8000", router))
}
