package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"packages/auth"
	"packages/models"
)

func Login(rw http.ResponseWriter, r *http.Request) {
	userByRequest := models.User{}
	if err := json.NewDecoder(r.Body).Decode(&userByRequest); err != nil {
		return
	}
	if user, err := models.GetUserByEmail(userByRequest.Email); err != nil {
		return
	} else {
		if user.Email == userByRequest.Email && user.Password == userByRequest.Password {
			user.Password = ""
			token := auth.GenerateJWT(*user)
			respToken := models.ResponseToken{
				Token: token,
			}

			models.SendData(rw, respToken)
		} else {
			fmt.Fprintln(rw, "Wrong email or password!")
		}
	}
}
