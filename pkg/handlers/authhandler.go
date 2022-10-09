package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//mySigningKey can be sent from secure place as env variable this has to match the variable sent by the client
var mySigningKey = []byte("ultimateStarAjith")

//AuthHandler handles all the operations related to the JWT Tokens
type AuthHandler struct {
}

//NewMiddlewareHandler returns a new instance of the middleware Handler
func NewMiddlewareHandler() *AuthHandler {
	return &AuthHandler{}
}

//GenerateToken generates a new JWT Token
func (handler *AuthHandler) GenerateToken(w http.ResponseWriter, r *http.Request) {
	tokenString, err := GenerateJWT()
	if err != nil {
		fmt.Println("error occured while generating the token string")
	}

	fmt.Fprintf(w, tokenString)
}

//GenerateJWT generates JWT Token
func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "Niroop"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return tokenString, nil
}
