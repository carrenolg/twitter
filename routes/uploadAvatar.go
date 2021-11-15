package routes

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/carrenolg/twitter/db"
	"github.com/carrenolg/twitter/models"
)

func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	// 1. Get image from URL
	image, header, err := r.FormFile("avatar")
	var ext = strings.Split(header.Filename, ".")[1]
	var path string = "uploads/avatars/" + IDUser + "." + ext

	// 2. Save image into the server
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		msg := fmt.Errorf("error during avatar uploading - %s", err)
		http.Error(w, msg.Error(), http.StatusInternalServerError)
		return
	}
	_, err = io.Copy(file, image)
	if err != nil {
		msg := fmt.Errorf("error during avatar uploading - %s", err)
		http.Error(w, msg.Error(), http.StatusInternalServerError)
		return
	}

	// 3. Update user model
	var user models.User
	var status bool
	user.Avatar = IDUser + "." + ext
	status, err = db.ModifyProfile(user, IDUser)
	if err != nil || status == false {
		msg := fmt.Errorf("error during avatar uploading - %s", err)
		http.Error(w, msg.Error(), http.StatusInternalServerError)
		return
	}

	// 4. Create response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
