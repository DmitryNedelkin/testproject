package rest_server

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"testproject/database"
)

var (
	mux sync.Mutex
)

func StartRestServer() {
	http.HandleFunc("/hello", handleMessage)

	log.Println("Starting rest server on :8090")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}
}

func handleMessage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getMessage(w)
	case http.MethodPost:
		addMessages(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func addMessages(w http.ResponseWriter, r *http.Request) {
	mux.Lock()
	defer mux.Unlock()

	var messagesInReq database.Messages
	if err := json.NewDecoder(r.Body).Decode(&messagesInReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.PutMessagesToDatabase(messagesInReq)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(database.AllMessages)
}

func getMessage(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(database.AllMessages)
}
