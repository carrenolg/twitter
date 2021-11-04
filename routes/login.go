package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/carrenolg/twitter/db"
	"github.com/carrenolg/twitter/jwt"
	"github.com/carrenolg/twitter/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	// user
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		msg := fmt.Errorf("input data error: %s", err)
		http.Error(w, msg.Error(), 400)
		return
	}
	// validations input data (email)
	if len(user.Email) == 0 {
		msg := fmt.Errorf("email field is required")
		http.Error(w, msg.Error(), 400)
		return
	}

	// try login
	document, loggedIn := db.TryLogin(user.Email, user.Password)
	if loggedIn == false {
		msg := fmt.Errorf("Invalid email or password")
		http.Error(w, msg.Error(), 400)
		return
	}

	// generate token
	jwtKey, err := jwt.GenerateJwt(document)
	if err != nil {
		msg := fmt.Errorf("error during token generation: %s", err)
		http.Error(w, msg.Error(), 400)
		return
	}

	// response
	resp := models.ResponseLogin{
		Toke: jwtKey,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	// create cookie
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
