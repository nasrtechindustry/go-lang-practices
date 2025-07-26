package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var PORT int = 8000

const TRIES int = 20

type User struct {
	Name    string   `json:"username"`
	Age     string   `json:"age"`
	Hobbies []string `json:"hobbies"`
}

func EntryPoint() {

	// creating the fall back mechanism

	http.HandleFunc("/user", welcome)

	for i := 1; i < TRIES; i++ {

		fmt.Println("Connecting Http server via PORT : " + strconv.Itoa(PORT))
		err := http.ListenAndServe(":"+strconv.Itoa(PORT), nil)
		time.Sleep(time.Second)

		if err != nil {
			PORT = PORT + i
			fmt.Println("The port was in use .Trying PORT :" + strconv.Itoa(PORT))
		}

	}

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
