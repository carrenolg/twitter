package routes

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/carrenolg/twitter/db"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	// 1. Get params from URL
	typeUser := r.URL.Query().Get("type")
	inputPage := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	// 2. Validate input parameter (page)
	page, err := strconv.Atoi(inputPage)
	if err != nil {
		msg := errors.New("page parameter is missing")
		http.Error(w, msg.Error(), http.StatusBadRequest)
		return
	}

	// 3. Call to RaadAllUsers()
	users, status := db.ReadAllUsers(IDUser, int64(page), search, typeUser)
	if status != true {
		msg := errors.New("error during users reading")
		http.Error(w, msg.Error(), http.StatusInternalServerError)
		return
	}

	// 4. Create response Http
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)

}
