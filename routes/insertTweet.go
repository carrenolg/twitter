package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/carrenolg/twitter/db"
	"github.com/carrenolg/twitter/models"
)

func InsertTweet(w http.ResponseWriter, r *http.Request) {
	// decode request
	var tweet models.Tweet
	err := json.NewDecoder(r.Body).Decode(&tweet)
	if err != nil {
		msg := fmt.Errorf("input data error: %s", err)
		http.Error(w, msg.Error(), 400)
		return
	}

	// create new tweet model
	newTweet := models.Tweet{
		UserId:   IDUser,
		Message:  tweet.Message,
		DateTime: time.Now(),
	}

	// call insert func
	_, status, err := db.InsertTweet(newTweet)
	if err != nil {
		msg := fmt.Errorf("error during tweet inserting: %s", err)
		http.Error(w, msg.Error(), 400)
		return
	}
	if status == false {
		msg := fmt.Errorf("can't insert tweet")
		http.Error(w, msg.Error(), 400)
		return
	}

	// create response
	w.WriteHeader(http.StatusCreated)
}
