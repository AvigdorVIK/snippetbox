package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("privet is snapchat"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("storinka vidobrajenya"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("storinka stvorenya novoi statti"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("zapusk servera na port :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
