package api

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Name    string   `json:"username"`
	Age     string   `json:"age"`
	Hobbies []string `json:"hobbies"`
}

func API() {

	http.HandleFunc("/", welcome)

	http.ListenAndServe(":8000", nil)

}

func welcome(w http.ResponseWriter, r *http.Request) {

	user := User{
		Name:    "Juma",
		Age:     "kaka",
		Hobbies: []string{"dancing", "cooking"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
