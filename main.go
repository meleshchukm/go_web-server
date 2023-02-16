package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		usersFile, err := ioutil.ReadFile("users.json")
		if err != nil {
			http.Error(w, "Error during reading users file", http.StatusInternalServerError)
			return
		}

		var users []User
		err = json.Unmarshal(usersFile, &users)
		if err != nil {
			http.Error(w, "Error during decoding users file", http.StatusInternalServerError)
			return
		}

		respBody, err := json.Marshal(users)
		if err != nil {
			http.Error(w, "Error during encoding response body", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(respBody)
	})

	fmt.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
