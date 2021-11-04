package db

import (
	"github.com/carrenolg/twitter/models"
	"golang.org/x/crypto/bcrypt"
)

func TryLogin(email string, password string) (models.User, bool) {
	// check if user is already exist
	user, found, _ := CheckUserExist(email)
	if found == false {
		return user, false
	}
	inputPassword := []byte(password)
	dbPassword := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(dbPassword, inputPassword)
	if err != nil {
		return user, false
	}
	return user, true
}
