package main

import (
	"log"
	"net/http"
)

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/paste", showpaste)
	mux.HandleFunc("/paste/new",  createpaste)

	log.Println("Starting server on localhost:8080")
	err := http.ListenAndServe("8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}