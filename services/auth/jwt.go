package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/MohamedMosalm/GO-E-Commerce/utils"
	"github.com/golang-jwt/jwt/v5"
)

var (
	JWTSECRETKEY = []byte(os.Getenv("JWTSECRET"))
)

func CreateJWT(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":     email,
		"expiresAt": time.Now().Add(24 * time.Hour).Unix(),
	})
	return token.SignedString(JWTSECRETKEY)
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return JWTSECRETKEY, nil
	})
}

func AuthMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("Unauthorized"))
			return
		}

		_, err = ValidateJWT(cookie.Value)

		if err != nil {
			utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("Unauthorized"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
