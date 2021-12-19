package handlers

import (
	"encoding/json"
	"net/http"
	"packages/models"
)

func NewUser(rw http.ResponseWriter, r *http.Request) {
	user := models.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		models.SendUnprocessableEntity(rw)
	} else {
		user.NewUser()
		models.SendData(rw, user)
	}
}
