package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/williamnoble/goExercises/database/postgres/RestApi_Gorm_JWT_PG/models"

	u "github.com/williamnoble/RestApi_Gorm_JWT_PG/util"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Response(w, u.Message(false, "invalidRequest"))
		return
	}
	resp := models.Login(account.Email, account.Password)
	u.Response(w, resp)
}
