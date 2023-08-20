package controllers

import (
	"RESTful-API/models"
	"encoding/json"
	"net/http"
)

type PingController struct{}

func NewPingController() *PingController {
	return &PingController{}
}

func (c *PingController) Ping(w http.ResponseWriter, r *http.Request) {
	ping := models.Ping{Status: "pong"}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ping)
}
