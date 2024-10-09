package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func IsTokenExpired(tokenString string) (bool, error) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("dummy-key"), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false, errors.New("invalid token claims")
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return false, errors.New("expiration claim not found or invalid")
	}

	return time.Now().Unix() > int64(exp), nil
}

func ExtractGuestID(tokenString string) (string, error) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("dummy-key"), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	guestID, ok := claims["guestId"].(string)
	if !ok {
		return "", errors.New("guestId claim not found or invalid")
	}

	return guestID, nil
}
