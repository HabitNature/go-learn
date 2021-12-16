package common

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"time"
)
import jwtgo "github.com/dgrijalva/jwt-go"

// 参见 ：https://www.cnblogs.com/jianga/p/12487267.html
// 参见 ：https://www.jianshu.com/p/39b584b47260
// 参见 : https://www.cnblogs.com/CYD-self/p/13954661.html

const defaultKey string = "default_jwt_key"

type Claims struct {
	User     string `json:"user"`
	Password string `json:"password"`
	jwtgo.StandardClaims
}

func GenerateToken(user, pwd string) (string, error) {
	jwtConfig := viper.GetStringMapString("jwt")
	key := jwtConfig["key"]
	if key == "" {
		key = defaultKey
	}

	exp := time.Now().Add(60 * time.Minute)

	c := Claims{
		User:     user,
		Password: pwd,
		StandardClaims: jwtgo.StandardClaims{
			// 过期时间
			ExpiresAt: exp.Unix(),
			// token 发行人
			Issuer: "go-learn",
		},
	}

	tokenClaims := jwtgo.NewWithClaims(jwtgo.SigningMethodES256, c)
	token, err := tokenClaims.SignedString([]byte(key))

	if err != nil {
		return "", err
	}

	return token, nil
}

func ParseToken(token string) (*Claims, error) {
	jwtConfig := viper.GetStringMapString("jwt")
	key := jwtConfig["key"]
	if key == "" {
		key = defaultKey
	}

	c, err := jwtgo.ParseWithClaims(token, &Claims{}, func(t *jwtgo.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := c.Claims.(*Claims)
	if ok && c.Valid {
		return claims, nil
	}

	return nil, errors.New(fmt.Sprintf("解析token格式不正确 ： %v", claims))
}
