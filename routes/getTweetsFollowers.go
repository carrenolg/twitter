package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/carrenolg/twitter/db"
)

func GetTweetsFollowers(w http.ResponseWriter, r *http.Request) {
	// 1. get params from URL
	pageParameter := r.URL.Query().Get("page")
	if len(pageParameter) < 1 {
		msg := errors.New("page parameter is missing")
		http.Error(w, msg.Error(), http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(pageParameter)
	if err != nil {
		msg := fmt.Errorf("page parameter must be a number - error: %s", err)
		http.Error(w, msg.Error(), http.StatusBadRequest)
		return
	}

	// 2. Call to ReadTweetsFollowers()
	tweets, status := db.ReadTweetsFollowers(IDUser, page)
	if status == false {
		msg := errors.New("error during of the tweets followers reading")
		http.Error(w, msg.Error(), http.StatusBadRequest)
		return
	}

	// 3. Create response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tweets)
}
