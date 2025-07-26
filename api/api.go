package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Book struct {
	Name      string
	Publisher string
	Year      int
	Authors   []string
}

type books []Book

func Init() {
	r := chi.NewRouter()


	books := books{Book{Name: "End of Vibe coding", Authors: []string{"juma", "nyambari"}},
		Book{Name: "Imperative programming", Authors: []string{"bro code", "Rizy codes"}}}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root api call"))
	})

	r.Get("/books", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(books)
	})

	http.ListenAndServe(":3000", r)
}
