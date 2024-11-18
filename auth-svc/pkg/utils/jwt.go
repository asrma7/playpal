package utils

import (
	"errors"
	"github.com/asrma7/playpal/auth-svc/internal/models"
	"github.com/golang-jwt/jwt"
	"time"
)

type JWTWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type jwtClaims struct {
	jwt.StandardClaims
	Id    int64
	Email string
}

func (w *JWTWrapper) GenerateToken(user models.User) (signedToken string, err error) {
	claims := &jwtClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    w.Issuer,
			ExpiresAt: time.Now().Local().Add(time.Duration(w.ExpirationHours) * time.Hour).Unix(),
		},
		Id:    user.Id,
		Email: user.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte(w.SecretKey))

	if err != nil {
		return "", err
	}

	return token.SignedString([]byte(w.SecretKey))
}

func (w *JWTWrapper) ValidateToken(signedToken string) (claims *jwtClaims, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken, &jwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(w.SecretKey), nil
		})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwtClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("token has expired")
	}

	return claims, nil
}
