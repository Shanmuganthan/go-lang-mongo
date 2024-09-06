package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Shanmuganthan/go-lang-mongo/common"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		token, err := common.VerifyJWTToken(authHeader)

		if err != nil {

			log.Println("Token validation failed")
			http.Error(w, "Un Authorized", http.StatusUnauthorized)
			return
		}

		fmt.Printf("Token Claim %T", token)

		ctx := context.WithValue(r.Context(), "user_details", token)

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
