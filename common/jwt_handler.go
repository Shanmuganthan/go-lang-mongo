package common

import (
	"fmt"
	"os"
	"strings"

	"github.com/Shanmuganthan/go-lang-mongo/models"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTToken(claims models.UserClaims) (string, error) {

	secretKey := os.Getenv("SECRET_KET")

	fmt.Println("SECRET KEY ACCESSED FROM env File", secretKey)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))

	return tokenString, err

}

func VerifyJWTToken(authHeader string) (*models.UserClaims, error) {

	tokens := strings.Split(authHeader, " ")

	if len(tokens) != 2 || tokens[0] != "Bearer" {
		return nil, fmt.Errorf("invalid token format")
	}

	tokenString := tokens[1]

	claims := &models.UserClaims{}

	secretKey := []byte(os.Getenv("SECRET_KET"))

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		// Validate the signing method to avoid security vulnerabilities
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return secretKey, nil
	})

	// Check for parsing errors or invalid token
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %v", err)
	}

	// Verify the token's validity
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Return the claims extracted from the token
	return claims, nil

}
