package main

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	pb "github.com/whitesyn/shippy/user-service/proto/user"
)

var (
	key = []byte("YvCQKVAiqHN6RO4OwhCr129pC76kfyqY")
)

// CustomClaims type
type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

// Authable interface
type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

// TokenService structure
type TokenService struct {
}

// Decode JWT-token and return custom claims
func (srv *TokenService) Decode(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*CustomClaims)
	if ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

// Encode user metadata and return token
func (srv *TokenService) Encode(user *pb.User) (string, error) {
	expiresAt := time.Now().Add(time.Hour * 24).Unix()
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expiresAt,
			Issuer:    "go.micro.srv.user",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}
