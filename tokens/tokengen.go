package tokens

import (
	"ecommerce-golang/database"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

type SignedDetails struct {
	Email          string
	First_Name     string
	Last_Name      string
	Uid            string
	StandardClaims jwt.MapClaims
}

var SECRET_KEY = os.Getenv("secret_key")
var UserData *mongo.Collection = database.UserData(database.Client, "Users")

func TokenGenerator(email string, firstname string, lastname string, uid string) (signedToken string, signedRefreshToken string, err error) {
	claims := jwt.MapClaims{
		"Email":      email,
		"First_Name": firstname,
		"Uid":        uid,
		"ExpiresAt":  time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
	}

	refreshClaims := jwt.MapClaims{
		"Uid":       uid,
		"ExpiresAt": time.Now().Local().Add(time.Hour * 168).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		return "", "", err
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS384, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}

func ValidateToken() {

}

func UpdateAllTokens() {

}
