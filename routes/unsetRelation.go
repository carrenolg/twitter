package routes

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/carrenolg/twitter/db"
	"github.com/carrenolg/twitter/models"
)

func UnsetRelation(w http.ResponseWriter, r *http.Request) {
	// 1. Set up database
	id := r.URL.Query().Get("id")
	if len(id) < 1 {
		msg := errors.New("id parameter is missing")
		http.Error(w, msg.Error(), http.StatusBadRequest)
		return
	}

	// 2. Call to DeleteRelation()
	relation := models.Relation{
		Userid:           IDUser,
		UseridRelationed: id,
	}
	status, err := db.DeleteRelation(relation)
	if err != nil {
		msg := fmt.Errorf("error during relation deleting - %s", err)
		http.Error(w, msg.Error(), http.StatusInternalServerError)
		return
	}

	// 3. Check operation satatus
	if status == false {
		msg := errors.New("can't delete relation")
		http.Error(w, msg.Error(), http.StatusInternalServerError)
		return
	}

	// 4. Create response
	w.WriteHeader(http.StatusCreated)
}
