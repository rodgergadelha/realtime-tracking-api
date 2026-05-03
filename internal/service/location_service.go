package service

import (
	"fmt"
	"realtime-tracking/internal/manager"
	"realtime-tracking/internal/model"
)

type LocationService struct {
	manager *manager.ClientManager
}

func NewLocationService(m *manager.ClientManager) *LocationService {
	return &LocationService{manager: m}
}

func (s *LocationService) UpdateLocation(loc model.Location) {
	clients := s.manager.GetClients(loc.DriverID)

	for _, client := range clients {
		err := client.Conn.WriteJSON(loc)
		if err != nil {
			fmt.Println("error sending, removing client")
			client.Conn.Close()
			s.manager.RemoveClient(client)
		}
	}
}