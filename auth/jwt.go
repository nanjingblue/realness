package auth

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"realness/model"
	"time"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

type Claims struct {
	UserID int32
	jwt.StandardClaims
}

// ReleaseToken 生成token
func ReleaseToken(user model.User) (string, error) {
	// token 过期时间
	expirationTime := time.Now().Add(14 * 24 * time.Hour)
	claims := &Claims {
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			// 发放时间
			ExpiresAt: expirationTime.Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer: "realness",
			Subject: "token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseUserToken 解析 Token
func ParseUserToken(tokenString string) (*jwt.Token, *Claims, error)  {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
