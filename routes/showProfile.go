package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/carrenolg/twitter/db"
)

func ShowProfile(w http.ResponseWriter, r *http.Request) {
	// get Id parameter
	id := r.URL.Query().Get("id")
	if len(id) < 1 {
		err := fmt.Errorf("Id parameter is missing in the request")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// search user into database
	profile, err := db.SearchProfile(id)
	if err != nil {
		err := fmt.Errorf("Error during searching of profile, id:%s", id)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// create response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profile)
}
