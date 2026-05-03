package handler

import (
	"fmt"
	"net/http"
	"realtime-tracking/internal/manager"

	"github.com/gorilla/websocket"
)

type WSHandler struct {
	manager *manager.ClientManager
}

func NewWSHandler(m *manager.ClientManager) *WSHandler {
	return &WSHandler{manager: m}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (h *WSHandler) Handle(w http.ResponseWriter, r *http.Request) {
	driverID := r.URL.Query().Get("driverId")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	client := &manager.Client{
		Conn:     conn,
		DriverID: driverID,
	}

	h.manager.AddClient(client)

	fmt.Println("client connected:", driverID)

	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			fmt.Println("client disconnected:", driverID)
			conn.Close()
			h.manager.RemoveClient(client)
			break
		}
	}
}