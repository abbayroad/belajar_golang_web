package main

import (
	"golang_web/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	aboutHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("About"))
	}

	mux.HandleFunc("/", handler.HelloHandler)
	mux.HandleFunc("/profile", handler.ProfileHandler)
	mux.HandleFunc("/produk", handler.ProdukHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Contact"))
	})

	fileserver := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static",fileserver))

	log.Println("Starting web at 8000")
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}

