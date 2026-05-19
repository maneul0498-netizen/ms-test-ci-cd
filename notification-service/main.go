package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
}

func notify(w http.ResponseWriter, r *http.Request) {

	var user User

	json.NewDecoder(r.Body).Decode(&user)

	log.Println("sending notification to:", user.Name)

	w.Write([]byte("notification sent"))
}

func main() {

	http.HandleFunc("/notify", notify)

	log.Println("notification-service running on 8082")

	http.ListenAndServe(":8082", nil)
}
