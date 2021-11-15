package routes

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/carrenolg/twitter/db"
)

func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	// 1. get params from URL
	id := r.URL.Query().Get("id")
	if len(id) < 1 {
		msg := errors.New("id parameter is missing")
		http.Error(w, msg.Error(), http.StatusBadRequest)
		return
	}

	// 2. call to DeleteTweet()
	err := db.DeleteTweet(id, IDUser)
	if err != nil {
		msg := fmt.Errorf("error during tweet deleting - %s", err)
		http.Error(w, msg.Error(), http.StatusBadRequest)
		return
	}

	// 3. create response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
