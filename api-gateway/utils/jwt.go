package utils

import (
	"api-gateway/config"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
    UserID string `json:"user_id"`
    jwt.StandardClaims
}

func GenerateJWT(userID string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        UserID: userID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    jwtKey := []byte(config.GetConfig().JWTSecret)
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func ParseJWT(tokenString string) (*Claims, error) {
    claims := &Claims{}
    jwtKey := []byte(config.GetConfig().JWTSecret)
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil {
        return nil, err
    }

    if !token.Valid {
        return nil, errors.New("invalid token")
    }

    return claims, nil
}
