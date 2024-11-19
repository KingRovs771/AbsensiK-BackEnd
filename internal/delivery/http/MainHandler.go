package http

import (
	"encoding/json"
	"net/http"
)

type MainHandler struct{}



func NewMainHandler()  *MainHandler{
	return &MainHandler{}
}

func (h *MainHandler) GetDomain(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
		"message": "Server Berjalan dengan baik", 
	})
}
