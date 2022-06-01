package util

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

var (
	Secret = []byte("TikTok")
	// TokenExpireDuration = time.Hour * 2 过期时间
)

type JWTClaims struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"user_name"`
	jwt.StandardClaims
}

//生成token
func GenToken(userid int64, userName string) (string, error) {
	claims := JWTClaims{
		UserId:   userid,
		Username: userName,
		StandardClaims: jwt.StandardClaims{
			Issuer: "server",
			//ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),可用于设定token过期时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("TikTok"))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

//解析token
func ParsenToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

//验证token
func VerifyToken(tokenString string) (int64, error) {

	fmt.Printf("tokenString:%v", tokenString)

	if tokenString == "" {
		return int64(0), nil
	}
	claims, err := ParsenToken(tokenString)
	if err != nil {
		return int64(0), err
	}
	return claims.UserId, nil
}
