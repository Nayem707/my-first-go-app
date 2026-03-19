package main

import (
	"fmt"
	"my-go-app/routes"
	"net/http"
)

func main() {
    routes.SetupRoutes()
    fmt.Println("Server running at http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
