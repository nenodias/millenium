package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/nenodias/millenium/configs"
)

type AuthKey string

const AUTH_KEY AuthKey = "auth"

var secretKey = []byte("SecretYouShouldHide")

type Token struct {
	Token string `json:"token"`
}

func GenerateJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Audience:  "millenium",
		Subject:   "millenium",
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(60 * time.Minute).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Verify(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &jwt.StandardClaims{
			Audience:  claims["aud"].(string),
			Subject:   claims["sub"].(string),
			IssuedAt:  int64(claims["iat"].(float64)),
			ExpiresAt: int64(claims["exp"].(float64)),
		}, nil
	} else {
		return nil, fmt.Errorf("expired token")
	}
}

func Init() {
	secretKey = []byte(configs.GetEnv(configs.SERVER_SECRET, "secret"))
}
