package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
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

func SendRequest(logger *logrus.Logger, url string, method string, payload interface{}) (interface{}, error) {
	logger.Infof("Sending %s request to %s with payload: %+v\n", method, url, payload)

	client := &http.Client{}
	var req *http.Request
	var err error

	if payload != nil {
		// Encode payload as JSON
		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}

		// Create HTTP request
		req, err = http.NewRequest(method, url, bytes.NewBuffer(jsonPayload))
		if err != nil {
			return nil, err
		}
	} else {
		// Create HTTP request without a request body
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			return nil, err
		}
	}

	// Set request headers
	req.Header.Set("Content-Type", "application/json")

	// Send HTTP request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check for HTTP error
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, resp.Status)
	}

	// Decode response body
	var response interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
