package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
)

var JWT_KET = []byte(os.Getenv("JWT_KEY"))

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/*
			This paths can be reached without the need of a token
		*/
		publicPaths := []string{
			"/auth/register",
			"/auth/login",
			"/join/", //TODO find a way to fix this
		}

		for _, path := range publicPaths {
			fmt.Println("Verifying path")
			//Verifying if the url have someting from the publicPaths
			if strings.HasPrefix(r.URL.Path, path) {
				next.ServeHTTP(w, r)
				return
			}
		}
		fmt.Println("Valid Path")
		// Extract token from the Authorization header
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "No token provided", http.StatusUnauthorized)
			return
		}

		// Parse the JWT token using StandardClaims
		claims := &jwt.StandardClaims{}

		//The func is to validate the token with our secret key
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return JWT_KET, nil
		})

		if err != nil {
			log.Printf("Error parsing token: %v", err)
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Token is valid, continue with the request
		next.ServeHTTP(w, r)
	})
}
