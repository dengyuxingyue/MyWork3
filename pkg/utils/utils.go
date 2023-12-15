package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// 用于签名的字符串
var JWTsecret = []byte("dengyuxingyue")

type Claims struct {
	Id       uint   `json:"id"` //区分用户id
	UserName string `json:"user_name"`
	Password string `json:"password"`
	jwt.RegisteredClaims
	/*
	   Issuer string `json:"iss,omitempty"`
	   Subject string `json:"sub,omitempty"`
	   Audience ClaimStrings `json:"aud,omitempty"`
	   ExpiresAt *NumericDate `json:"exp,omitempty"`
	   NotBefore *NumericDate `json:"nbf,omitempty"`
	   IssuedAt *NumericDate `json:"iat,omitempty"`
	   ID string `json:"jti,omitempty"`
	*/
}

// 签发token
func GenerateToken(id uint, username, password string) (string, error) {
	claims := Claims{
		Id:       id,
		UserName: username,
		Password: password,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), //设置过期时间
			Issuer:    "todo_list",                                        //发行者
		},
	}
	// 生成token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 生成签名字符串
	return token.SignedString(JWTsecret)

}

// 验证token
// ParseToken 解析JWT
func ParseToken(tokenString string) (*Claims, error) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return JWTsecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*Claims); ok && token.Valid { // 校验token

		return claims, nil

	}

	return nil, errors.New("invalid token")

}
