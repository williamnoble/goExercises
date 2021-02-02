package controllers

import (
	"encoding/json"
	"net/http"

	u "github.com/williamnoble/RestApi_Gorm_JWT_PG/util"

	"github.com/williamnoble/RestApi_Gorm_JWT_PG/models"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(uint)
	contact := &models.Contact{}
	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Response(w, u.Message(false, "error whilst decoding body request"))
		return
	}

	contact.UserId = user
	resp := contact.Create()
	u.Response(w, resp)
}
