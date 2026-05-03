package main

import (
	"fmt"
	"net/http"
	"os"

	"realtime-tracking/internal/handler"
	"realtime-tracking/internal/manager"
	"realtime-tracking/internal/service"
)

func main() {
	manager := manager.NewClientManager()
	service := service.NewLocationService(manager)

	httpHandler := handler.NewHTTPHandler(service)
	wsHandler := handler.NewWSHandler(manager)

	http.HandleFunc("/location", httpHandler.Location)
	http.HandleFunc("/ws", wsHandler.Handle)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server running on :", port)
	http.ListenAndServe(":"+port, nil)
}