package routes

import (
	"errors"
	"log"
	"strings"

	"github.com/carrenolg/twitter/db"
	"github.com/carrenolg/twitter/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

// export variables
var Email string
var IdUser string

var JwtKey = GetKey()

func CheckToken(token string) (*models.Claim, bool, string, error) {
	myPass := []byte(JwtKey)
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		err := errors.New("invalid token format")
		return claims, false, string(""), err
	}
	token = strings.TrimSpace(splitToken[1])

	// parse token
	jwToken, err := jwt.ParseWithClaims(
		token,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return myPass, nil
		})

	// check relationship between email and token
	if err == nil {
		_, isExist, _ := db.CheckUserExist(claims.Email)
		if isExist {
			Email = claims.Email
			IdUser = claims.Id
		}
		return claims, isExist, IdUser, nil
	}

	// invalid token
	if !jwToken.Valid {
		err := errors.New("invalid token")
		return claims, false, string(""), err
	}

	return claims, false, string(""), err
}

func GetKey() string {
	var envs map[string]string
	envs, err := godotenv.Read("routes/.env")
	if err != nil {
		log.Printf("err: %s", err)
		log.Fatal("Error loading .env file")
	}
	key := envs["TOKEN_PASS"]
	return key
}
