package token

import (
	"encoding/json"
	"example-app/pkg/server"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Auth struct {
	AccessToken  string
	RefreshToken string
	Email        string
}

var (
	secretKey = []byte("secret-key")
	cache     = make(map[string]Auth)
)

func CreateToken(auth Auth) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"auth-token":    auth.AccessToken,
			"refresh-token": auth.RefreshToken,
			"exp":           time.Now().Add(time.Hour * 24).Unix(),
			"email":         auth.Email,
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	cache[tokenString] = auth

	return tokenString, nil
}

func GetAccessToken(tokenString string) (string, error) {
	var accessToken string

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return accessToken, err
	}

	if !token.Valid {
		return accessToken, fmt.Errorf("invalid token")
	}

	auth, ok := cache[tokenString]
	if !ok {
		return accessToken, fmt.Errorf("invalid token")
	}

	// TODO: use refresh token if access token expired

	return auth.AccessToken, nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var auth Auth
	json.NewDecoder(r.Body).Decode(&auth)

	if auth.AccessToken != "" && auth.RefreshToken != "" {
		tokenString, err := CreateToken(auth)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			// fmt.Errorf("auth info not found")
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, tokenString)
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid credentials")
	}
}

func RegisterHandlers() {
	server.AddRoute("/login", LoginHandler, "POST")
}
