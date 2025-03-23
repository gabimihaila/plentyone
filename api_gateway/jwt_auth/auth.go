package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"api_gateway/config"
	"api_gateway/logger"

	"github.com/golang-jwt/jwt/v5"
)

// Auth is a standard HTTP middleware for JWT authentication
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		configuration := config.LoadConfig()

		token := r.Header.Get("Authorization")

		authTokens := strings.Split(token, " ")

		if len(authTokens) != 2 {
			http.Error(w, "Invalid authorization Header ", http.StatusUnauthorized)
			logger.Error(errors.New("invalid authorization Header"))
			return
		}

		tokenString := authTokens[1]

		logger.Info("Token string e " + tokenString)

		var ok bool

		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok = token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(configuration.JWTSecret), nil
		})

		if err != nil {
			http.Error(w, `{"error": "Invalid token"}`, http.StatusUnauthorized)
			logger.Error(err)
			return
		}

		logger.Info("Trece prin auth cu ce face: ")

		// Call the next handler if authentication is successful
		next.ServeHTTP(w, r)
	})
}
