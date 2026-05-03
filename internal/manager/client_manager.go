package manager

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	DriverID string
}

type ClientManager struct {
	clients map[string][]*Client
	mu      sync.RWMutex
}

func NewClientManager() *ClientManager {
	return &ClientManager{
		clients: make(map[string][]*Client),
	}
}

func (m *ClientManager) AddClient(c *Client) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.clients[c.DriverID] = append(m.clients[c.DriverID], c)
}

func (m *ClientManager) RemoveClient(c *Client) {
	m.mu.Lock()
	defer m.mu.Unlock()

	clients := m.clients[c.DriverID]

	for i, client := range clients {
		if client == c {
			m.clients[c.DriverID] = append(clients[:i], clients[i+1:]...)
			break
		}
	}

	if len(m.clients[c.DriverID]) == 0 {
		delete(m.clients, c.DriverID)
	}
}

func (m *ClientManager) GetClients(driverID string) []*Client {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.clients[driverID]
}