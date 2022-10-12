package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/configs"
	"github.com/models"
)

func GetCreds(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var credentials []models.UserCredentials
	configs.DataBase.Preload("credentials").Find(&credentials)
	json.NewEncoder(w).Encode(credentials)
}
