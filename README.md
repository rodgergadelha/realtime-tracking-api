# Real-Time Tracking API

API de rastreamento em tempo real construída com Go. Recebe localizações via HTTP e as envia instantaneamente para clientes conectados via WebSocket.

---

## Produção

```text
https://realtime-tracking-api.onrender.com
```

---

## Endpoints

### Enviar localização

```http
POST https://realtime-tracking-api.onrender.com/location
```

```json
{
  "driverId": "123",
  "lat": -3.73,
  "lng": -38.52
}
```

---

### Conectar via WebSocket

```http
GET wss://realtime-tracking-api.onrender.com/ws?driverId=123
```

---

## Como funciona

1. O motorista envia a localização via HTTP
2. O servidor recebe e processa
3. Clientes conectados recebem o update em tempo real via WebSocket

---

## Características

* Comunicação em tempo real (WebSocket)
* Múltiplos clientes por motorista
* Armazenamento em memória

---

## Rodar localmente

```bash
go run cmd/main.go
```

```text
Server running on : 8080
```

---

## Uso

* Rastreamento de entregas
* Monitoramento em tempo real
* Simulações de localização

---