package main

import (
	"net/http"
	"log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, _  *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	port := ":3000"
	log.Printf("Starting server on port %s/n", port)
	http.ListenAndServe(":3000", r)
}
