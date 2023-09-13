package main

import (
	// "fmt"
	"fmt"
	"log"
	"net/http"
	// "net/http"
	"tps/constants"
	"tps/routes"

	// "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = []byte("https://www.postman.com/spaceflight-pilot-19131968/workspace/jwt")

// Replace with your secret key
func createToken() (string, error) {
	claims := jwt.MapClaims{
		// Token expires in 24 hours
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
func handler(w http.ResponseWriter, r *http.Request) {
	tokenString, err := createToken()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "JWT Token: %s", tokenString)
}
func main() {
	router := gin.Default()
	routes.AppRoutes(router)
	// http.HandleFunc("/jwt", handler)
	// http.ListenAndServe(constants.Port, nil)
	log.Fatal(router.Run(constants.Port))
}
