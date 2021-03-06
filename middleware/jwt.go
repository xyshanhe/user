package middleware

import (
	"User/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("share_token")

type Claims struct {
	Userid uint
	jwt.StandardClaims
}

func ReleastToken(user model.User) (string, error) {

	//过期时间
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := &Claims{
		Userid: uint(user.Id),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),     //发放时间
			Issuer:    "share",               //签发人
			Subject:   "user token",
		},
	}

	// 使用指定的签名方法(hash)创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParesToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(token *jwt.Token) (i interface{}, err error) {
			return jwtKey, nil
		})
	return token, claims, err
}
