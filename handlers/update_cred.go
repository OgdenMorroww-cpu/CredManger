package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/configs"
	"github.com/models"
)

func UpdateCred(w http.ResponseWriter, r *http.Request) {
	var updatedCredentials models.UserCredentials
	json.NewDecoder(r.Body).Decode(&updatedCredentials)
	configs.DataBase.Save(&updatedCredentials)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedCredentials)
}
