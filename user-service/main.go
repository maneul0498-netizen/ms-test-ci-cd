package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
}

func createUser(w http.ResponseWriter, r *http.Request) {

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("user created:", user.Name)

	body, _ := json.Marshal(user)

	_, err = http.Post(
		"http://notification-service:8082/notify",
		"application/json",
		bytes.NewBuffer(body),
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	w.Write([]byte("user created"))
}

func main() {

	http.HandleFunc("/users", createUser)

	log.Println("user-service running on 8080")

	http.ListenAndServe(":8081", nil)
}
