package tokens

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTCustomClaims struct {
	Id    uint
	Env   string
	Email string
	jwt.StandardClaims
}

type JWTToken struct {
	Id      uint
	Name    string
	Email   string
	Env     string
	Version string
	RealId  uint
}

type Toker struct {
	Env    string
	Secret string
	Expiry time.Duration
}

func New(env, secret string, expiry time.Duration) Toker {
	return Toker{
		Env:    env,
		Secret: secret,
		Expiry: expiry,
	}
}

func (t Toker) Create(id uint, name, lastName, email string) (string, error) {
	claims := JWTCustomClaims{
		Id:    id,
		Email: email,
		Env:   t.Env,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(t.Expiry).Unix(),
			Issuer:    name + " " + lastName,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(t.Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

var ErrInvalidToken = errors.New("invalid token")
var ErrParsingToken = errors.New("failed to parse token")

func (t Toker) Get(tokenString string) (JWTToken, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.Secret), nil
	})
	if err != nil {
		return JWTToken{}, err
	}

	if claims, ok := token.Claims.(*JWTCustomClaims); ok {
		if !token.Valid {
			return JWTToken{}, ErrInvalidToken
		}
		return JWTToken{
			Id:    claims.Id,
			Name:  claims.StandardClaims.Issuer,
			Email: claims.Email,
			Env:   claims.Env,
		}, nil
	}
	return JWTToken{}, ErrParsingToken
}
