package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/configs"
	"github.com/gorilla/mux"
	"github.com/models"
)

func GetSingleCred(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputCredentialsID := params["credentialID"]

	var credentials models.UserCredentials
	configs.DataBase.Preload("Credentials").First(&credentials, inputCredentialsID)
	json.NewEncoder(w).Encode(credentials)
}
