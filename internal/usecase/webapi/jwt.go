package webapi

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dimk00z/go-musthave-diploma/internal/usecase"
)

type JWTClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

// https://codewithmukesh.com/blog/jwt-authentication-in-golang/
// https://seefnasrul.medium.com/create-your-first-go-rest-api-with-jwt-authentication-in-gin-framework-dbe5bda72817

func (g *GopherMartWebAPI) GenerateToken(userID string) (string, error) {
	expirationTime := time.Now().Add(time.Duration(g.cfg.Security.TokenHourLifespan) * time.Hour)
	claims := &JWTClaim{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(g.cfg.Security.SecretKey))

	return tokenString, err

}

func (g *GopherMartWebAPI) CheckToken(tokenString string) error {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(g.cfg.Security.SecretKey), nil
	})
	return err

}

func (g *GopherMartWebAPI) ParseToken(tokenString string) (userID string, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("GopherMartRepoWebAPI - ExtractTokenID - r.Builder: %v", token.Header["alg"])
		}
		return []byte(g.cfg.Security.SecretKey), nil
	})
	if err != nil {
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID, ok = claims["user_id"].(string)
		if ok {
			return
		}
	}
	return "", usecase.ErrJWT
}
