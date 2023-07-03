package http

import (
	"errors"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

func JWTAuth(original func(w http.ResponseWriter, r *http.Request),
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the value of the "Authorization" header from the request
		authHeader := r.Header["Authorization"]
		if authHeader == nil {
			// If the header is missing, respond with "not authorized" and HTTP status code 401 (Unauthorized)
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}

		// Split the "Authorization" header value into two parts: the scheme ("Bearer") and the token string
		authHeaderParts := strings.Split(authHeader[0], " ")
		// If the header value doesn't have two parts or the scheme is not "Bearer",
		// respond with "not authorized" and HTTP status code 401 (Unauthorized)
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}

		// Validate the incoming token by calling the validateToken function
		if validateToken(authHeaderParts[1]) {
			// If the token is valid, call the original handler function with the provided response writer and request
			original(w, r)
		} else {
			// If the token is not valid, respond with "not authorized" and HTTP status code 401 (Unauthorized)
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}
	}
}

func validateToken(accessToken string) bool {
	var mySigningKey = []byte("missionimpossible")
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// If the token's signing method is not HMAC, return an error indicating that the auth token could not be validated
			return nil, errors.New("could not validate auth token")
		}

		return mySigningKey, nil
	})

	if err != nil {
		// If there was an error parsing the token, return false indicating the token is not valid
		return false
	}

	return token.Valid
}
