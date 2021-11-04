package jwt

import (
	"log"
	"time"

	"github.com/carrenolg/twitter/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

var JwtKey = GetKey()

func GenerateJwt(user models.User) (string, error) {
	myKey := []byte(JwtKey)

	payload := jwt.MapClaims{
		"email":      user.Email,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"date_birth": user.DateBirth,
		"biography":  user.Biography,
		"location":   user.Location,
		"website":    user.WebSite,
		"_id":        user.ID.Hex(),
		"exp":        time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	strToken, err := token.SignedString(myKey)
	if err != nil {
		return strToken, err
	}
	return strToken, nil
}

func GetKey() string {
	var envs map[string]string
	envs, err := godotenv.Read("jwt/.env")
	if err != nil {
		log.Printf("err: %s", err)
		log.Fatal("Error loading .env file")
	}
	key := envs["TOKEN_PASS"]
	return key
}
