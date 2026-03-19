package routes

import (
	"my-go-app/handlers"
	"net/http"
)

func SetupRoutes() {
    http.HandleFunc("/users", handlers.GetUsers)
}
