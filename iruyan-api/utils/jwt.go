package utils

import (
	"iruyan-api/config"
	"iruyan-api/pkg/errdefs"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(iruyanID string) (string, error) {
	if len(config.JWTSecret) == 0 {
		return "", errdefs.ErrJWTSecretNotSet
	}

	claims := jwt.MapClaims{
		"iruyanId": iruyanID,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	log.Printf("JWT TOKEN: %+v", token)

	signedToken, err := token.SignedString([]byte(config.JWTSecret))
	if err != nil {
		return "", errdefs.ErrJWTSignToken
	}

	return signedToken, nil
}

func ParseJWT(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return nil, nil, errdefs.ErrJWTSecretNotSet
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errdefs.ErrJWTInvalidToken
		}
		return []byte(secretKey), nil
	})
	if err != nil || !token.Valid {
		return nil, nil, errdefs.ErrJWTParseFailure
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, nil, errdefs.ErrJWTInvalidToken
	}

	return token, claims, nil
}
