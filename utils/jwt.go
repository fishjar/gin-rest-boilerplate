package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fishjar/gin-rest-boilerplate/config"
	"github.com/fishjar/gin-rest-boilerplate/model"
)

// JWTClaims JWT加密的结构体
type JWTClaims struct {
	AuthID string `json:"aid" binding:"required"`
	jwt.StandardClaims
}

// MakeToken 创建JWT的TOKEN
func MakeToken(user *model.UserJWT) (string, error) {

	signKey := config.JWTSignKey                                  // JWT加密用的密钥
	expiresAt := time.Duration(config.JWTExpiresAt) * time.Minute // JWT过期时间，分钟为单位
	mySigningKey := []byte(signKey)                               // 密钥格式转换

	claims := JWTClaims{
		user.AuthID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiresAt).Unix(), // 过期时间
			IssuedAt:  time.Now().Unix(),                // 签发时间
			Issuer:    "gin",                            // 签发主体
			Subject:   user.UserID,                      // 主题
			// Id:        user.UserID,                      // 编号是JWT的唯一标识
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(mySigningKey)
}

// ParseToken 解析JWT的TOKEN
func ParseToken(tokenString string) (*JWTClaims, error) {

	signKey := config.JWTSignKey    // JWT加密用的密钥
	mySigningKey := []byte(signKey) // 密钥格式转换

	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	// 验证成功
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	// 验证失败
	return nil, errors.New("验证失败")
}

// 测试JWT功能是否正常
func init() {
	// 测试生成token
	token, err := MakeToken(&model.UserJWT{
		AuthID: "123",
		UserID: "456",
	})
	if err != nil {
		fmt.Println(err)
		panic("JWT生成token出错")
	}

	// 测试解析token
	claims, err := ParseToken(token)
	if err != nil {
		fmt.Println(err)
		panic("JWT解析token出错")
	}

	fmt.Println("JWT正常----------------")
	fmt.Println("token:", token)
	fmt.Println("claims:", claims)
}
