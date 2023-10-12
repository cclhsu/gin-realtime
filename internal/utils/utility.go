package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

func GetUUIDFromToken(tokenString string) (string, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, nil)
	if err != nil {
		return "", fmt.Errorf("failed to parse JWT token: %v", err)
	}

	// Verify the token signature
	if !token.Valid {
		return "", fmt.Errorf("invalid JWT token")
	}

	// Extract the claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("failed to extract JWT claims")
	}

	// Extract the UUID from the claims
	UUID, ok := claims["UUID"].(string)
	if !ok {
		return "", fmt.Errorf("UUID not found in JWT claims")
	}

	return UUID, nil
}
