package main

import (
//     "fmt"
    "net/http"
    "encoding/json"
    "log"
    "sync"
)

type Message struct {
    Index int32 `json:"index"`
    Id string `json:"id"`
    Guid string `json:"guid"`
    IsActive bool `json:"isActive"`
}

type Messages []Message

var (
    messages = []Messages{}
    mux sync.Mutex
)

func main() {
    http.HandleFunc("/hello", handleMessage)

    log.Println("Starting server on :8090")
    if err := http.ListenAndServe(":8090", nil); err != nil {
        log.Fatal(err)
    }
}

func handleMessage(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
        case http.MethodGet:
            getMessage(w, r)
        case http.MethodPost:
            addMessages(w, r)
        default:
            w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

func addMessages(w http.ResponseWriter, r *http.Request) {
    mux.Lock()
    defer mux.Unlock()

    var messagesInReq Messages
    if err := json.NewDecoder(r.Body).Decode(&messagesInReq); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    messages = append(messages, messagesInReq)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(messages)
}

func getMessage(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(messages)
}