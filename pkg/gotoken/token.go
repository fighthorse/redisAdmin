package gotoken

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	LoginSecret = "dasEgsSgfGJOIJngsdakSeqdsa"
)

func CreateToken(uid string, secret string) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

// uid
func ParseToken(token string, secret string) (string, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}
	uid := claim.Claims.(jwt.MapClaims)["uid"].(string)
	return uid, nil
}
