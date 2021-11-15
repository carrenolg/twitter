package routes

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/carrenolg/twitter/db"
)

func GetBanner(w http.ResponseWriter, r *http.Request) {
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
		msg := fmt.Errorf("error during banner getting - %s", err)
		http.Error(w, msg.Error(), http.StatusInternalServerError)
		return
	}

	// 3. Open the file banner
	path := "uploads/banners/"
	var bannerFile *os.File
	bannerFile, err = os.Open(path + userModel.Banner)
	if err != nil {
		msg := fmt.Errorf("error during banner getting - %s", err)
		http.Error(w, msg.Error(), http.StatusInternalServerError)
		return
	}

	// 4. Copy file to Http response
	_, err = io.Copy(w, bannerFile)
	if err != nil {
		msg := fmt.Errorf("error during banner getting - %s", err)
		http.Error(w, msg.Error(), http.StatusInternalServerError)
		return
	}

}
