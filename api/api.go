package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const PORT int = 9090

type User struct {
	Name    string   `json:"username"`
	Age     string   `json:"age"`
	Hobbies []string `json:"hobbies"`
}

func API() {

	http.HandleFunc("/user", welcome)
	fmt.Printf("The server is listening on port %v via http://localhost:%v"  , PORT  ,PORT)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(PORT), nil))

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
