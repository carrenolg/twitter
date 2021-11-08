package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/carrenolg/twitter/db"
)

func ReadTweets(w http.ResponseWriter, r *http.Request) {
	// 1. get params from URL
	id := r.URL.Query().Get("id")
	if len(id) < 1 {
		msg := errors.New("id parameter is missing")
		http.Error(w, msg.Error(), http.StatusBadRequest)
		return
	}
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

	// 2. call to readTweets function
	result, status := db.ReadTweets(id, int64(page))
	if status == false {
		msg := errors.New("error during tweets reading")
		http.Error(w, msg.Error(), http.StatusBadRequest)
		return
	}

	// 3. create response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}
