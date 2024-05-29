package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type NewJwtClaim struct {
	UserIdtentify interface{}
	jwt.RegisteredClaims
}

type JWT struct {
	SigningKey string
	ExpireHour int
}

func (j *JWT) Obtaining(u interface{}) (string, error) {
	claims := NewJwtClaim{
		u,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(j.ExpireHour) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(j.SigningKey))
	return ss, err
}

func (j *JWT) Authenticating(s string) (interface{}, int, error) {
	token, err := jwt.ParseWithClaims(
		s,
		&NewJwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SigningKey), nil
		},
	)

	if token.Valid {
		if claims, ok := token.Claims.(*NewJwtClaim); ok && token.Valid {
			return claims, http.StatusOK, nil
		}
	}
	if errors.Is(err, jwt.ErrTokenMalformed) {
		return nil, http.StatusBadRequest, errors.New("token解析失败")
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		return nil, http.StatusForbidden, errors.New("无效的token")
	} else {
		return nil, http.StatusBadRequest, err
	}
}
