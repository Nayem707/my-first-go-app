package handlers

import (
	"encoding/json"
	"my-go-app/models"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
    users := []models.User{
        {ID: 1, Name: "Nayem", Email: "nayem@example.com"},
        {ID: 2, Name: "Alex", Email: "alex@example.com"},
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}
