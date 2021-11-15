package routes

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/carrenolg/twitter/db"
	"github.com/carrenolg/twitter/models"
)

func CheckRelation(w http.ResponseWriter, r *http.Request) {
	// 1. Get parms from URL
	id := r.URL.Query().Get("id")
	if len(id) < 1 {
		msg := errors.New("id parameter is missing")
		http.Error(w, msg.Error(), http.StatusBadRequest)
		return
	}

	// 2. Call to CheckRelation()
	relation := models.Relation{
		Userid:           IDUser,
		UseridRelationed: id,
	}
	var stateRes models.StateRelation
	state, err := db.CheckRelation(relation)
	if err != nil || state == false {
		stateRes.State = false
	} else {
		stateRes.State = true
	}

	// 3. Create response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stateRes)
}
