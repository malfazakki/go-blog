package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/malfazakki/go-blog/models"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

// GenerateToken creates a new JWT token for a user
func GenerateToken(user *models.User) (string, error) {
	// Set token expiration to 24 hours
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create claims with user ID and expiration time
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     expirationTime.Unix(),
	}

	// Create claims with user ID and expiration time
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken checks if a token is valid and returns the user ID
func ValidateToken(tokenString string) (uint, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return jwtKey, nil
	})

	if err != nil {
		return 0, err
	}

	// Check if token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check if token is expired
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return 0, errors.New("token expired")
		}

		// Get user ID from claims
		userID := uint(claims["user_id"].(float64))
		return userID, nil
	}

	return 0, errors.New("invalid token")
}
