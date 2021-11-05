package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/carrenolg/twitter/db"
	"github.com/carrenolg/twitter/models"
)

func ModifyProfile(w http.ResponseWriter, r *http.Request) {
	// mapping user model
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		msg := fmt.Errorf("input data error: %s", err)
		http.Error(w, msg.Error(), 400)
		return
	}

	// update profile
	var status bool
	status, err = db.ModifyProfile(user, IDUser)
	if err != nil {
		err := fmt.Errorf("Error during updating of profile - Error:%s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if status == false {
		err := fmt.Errorf("Can't update profile - idUser:%s", IDUser)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// create response
	w.WriteHeader(http.StatusCreated)
}
