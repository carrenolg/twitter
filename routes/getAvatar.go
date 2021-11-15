package routes

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/carrenolg/twitter/db"
)

func GetAvatar(w http.ResponseWriter, r *http.Request) {
	// 1. get params from URL
	id := r.URL.Query().Get("id")
	if len(id) < 1 {
		msg := errors.New("id parameter is missing")
		http.Error(w, msg.Error(), http.StatusBadRequest)
		return
	}

	// 2. Get user model from db
	userModel, err := db.SearchProfile(id)
	if err != nil {
		msg := fmt.Errorf("error during avatar getting - %s", err)
		http.Error(w, msg.Error(), http.StatusInternalServerError)
		return
	}

	// 3. Open the file avatar
	path := "uploads/avatars/"
	var avatarFile *os.File
	avatarFile, err = os.Open(path + userModel.Avatar)
	if err != nil {
		msg := fmt.Errorf("error during avatar getting - %s", err)
		http.Error(w, msg.Error(), http.StatusInternalServerError)
		return
	}

	// 4. Copy file to Http response
	_, err = io.Copy(w, avatarFile)
	if err != nil {
		msg := fmt.Errorf("error during avatar getting - %s", err)
		http.Error(w, msg.Error(), http.StatusInternalServerError)
		return
	}

}
