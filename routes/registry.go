package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/carrenolg/twitter/db"
	"github.com/carrenolg/twitter/models"
)

func Registry(w http.ResponseWriter, r *http.Request) {
	// check json decode (input model)
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		msg := fmt.Errorf("input data error: %s", err)
		http.Error(w, msg.Error(), 400)
		return
	}

	// validations input data (email, password)
	if len(user.Email) == 0 {
		msg := fmt.Errorf("email field is empty - UserId:%s", user.ID)
		http.Error(w, msg.Error(), 400)
		return
	}
	if len(user.Password) < 6 {
		msg := fmt.Errorf("password length is invalid (min. 6 characters) - UserId:%s", user.ID)
		http.Error(w, msg.Error(), 400)
		return
	}

	// check user exist (on database)
	_, find, _ := db.CheckUserExist(user.Email)
	if find == true {
		msg := fmt.Errorf("user already exist: %s", user.Email)
		http.Error(w, msg.Error(), 400)
		return
	}

	// check insert registry
	_, status, err := db.InsertRegistry(user)
	if status == false {
		msg := fmt.Errorf("failed insert user: %s", err.Error())
		http.Error(w, msg.Error(), 400)
		return
	}
}
