package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/configs"
	"github.com/models"
)

func CreatCred(w http.ResponseWriter, r *http.Request) {
	var credentials models.UserCredentials
	json.NewDecoder(r.Body).Decode(&credentials)

	configs.DataBase.Create(&credentials)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(credentials)
}
