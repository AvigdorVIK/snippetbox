package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("privet is snapchat"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	log.Println("zapusk servera na port ...")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
