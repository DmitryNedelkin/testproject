package rest_server

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"testproject/database"
)

func TestAddMessage(t *testing.T) {

	log.Println("start AddMessageTest")
	server := httptest.NewServer(http.HandlerFunc(handleMessage))
	defer server.Close()

	message := database.Message{
		Index:    5,
		Id:       "668942917a79d4ec35a3717c",
		Guid:     "59842498-b8cf-4ed5-8727-0e7fa69e119d",
		IsActive: true}

	body, _ := json.Marshal(message)

	resp, err := http.Post(server.URL+"/hello", "application/json", bytes.NewBuffer(body))

	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
	}

	var createdMessage database.Message
	if err := json.NewDecoder(resp.Body).Decode(&createdMessage); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if createdMessage.Guid != message.Guid ||
		createdMessage.Id != message.Guid ||
		createdMessage.Index != message.Index ||
		createdMessage.IsActive != message.IsActive {
		t.Fatalf("Expected message %+v, got %+v", message, createdMessage)
	}
}
