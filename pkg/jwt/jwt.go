package jwt

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/okh8609/gin_blog/global"
)

type Claims struct {
	UUID string `json:"uuid"`
	jwt.StandardClaims
}

func GetJWTKey() []byte {
	return []byte(global.Auth.JWTkey)
}

func GenerateJWTToken(uuid string) string {
	iat := time.Now().UTC()
	exp := iat.Add(time.Duration(global.Auth.JWTexp * int(time.Second)))

	claim := Claims{
		UUID: uuid,
		StandardClaims: jwt.StandardClaims{
			Issuer:    global.Auth.JWTiss, // 簽發者
			IssuedAt:  iat.Unix(),         // 簽發時間
			NotBefore: iat.Unix(),         // 生效時間
			ExpiresAt: exp.Unix(),         // 過期時間
			// Id:        "",                 // JWT的唯一標識
			// Audience:  "",                 // 接受JWT者
			// Subject:   "",                 // 主題
		},
	}

	jwt_token := jwt.NewWithClaims(jwt.SigningMethodHS512, claim)
	token, err := jwt_token.SignedString(GetJWTKey())
	if err == nil {
		return token
	} else {
		global.MyLogger.Errorf(context.Background(), "# GenerateJWTToken Error: %v", err)
		return ""
	}
}

func VerifyJWTToken(token string) (*Claims, error) {
	jwt_token, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) { return GetJWTKey(), nil })
	if err != nil {
		return nil, err
	}
	if jwt_token != nil {
		claims, ok := jwt_token.Claims.(*Claims)
		if ok && jwt_token.Valid {
			return claims, nil
		}
	}
	return nil, err
}
