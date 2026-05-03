package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"realtime-tracking/internal/model"
	"realtime-tracking/internal/service"
)

type HTTPHandler struct {
	service *service.LocationService
}

func NewHTTPHandler(s *service.LocationService) *HTTPHandler {
	return &HTTPHandler{service: s}
}

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func (h *HTTPHandler) Location(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)

	body, _ := io.ReadAll(r.Body)

	var loc model.Location
	json.Unmarshal(body, &loc)

	h.service.UpdateLocation(loc)

	w.WriteHeader(http.StatusOK)
}