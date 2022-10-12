package handlers

import (
	"net/http"
	"strconv"

	"github.com/configs"
	"github.com/gorilla/mux"
	"github.com/models"
)

func DeleteCreds(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	inputCredentialsID := params["credentialId"]

	id64, _ := strconv.ParseUint(inputCredentialsID, 10, 64)

	idToDelete := uint(id64)

	configs.DataBase.Where("credential_id = ?", idToDelete).Delete(&models.CredentialsIDs{})
	configs.DataBase.Where("credential_id = ?", idToDelete).Delete(&models.UserCredentials{})

	w.WriteHeader(http.StatusNoContent)
}
