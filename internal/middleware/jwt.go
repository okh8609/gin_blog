package middleware

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/okh8609/gin_blog/pkg/app"
	"github.com/okh8609/gin_blog/pkg/errcode"

	_jwt "github.com/okh8609/gin_blog/pkg/jwt"
)

func JWTMiddleware(c *gin.Context) {
	ecode := errcode.Success
	token := c.GetHeader("Authorization")
	if token == "" {
		ecode = errcode.InvalidParams
	} else {
		if strings.Contains(token, "Bearer ") { // 去掉 Bearer
			_token := strings.Split(token, "Bearer ")
			token = strings.TrimSpace(_token[1])
		}

		claim, err := _jwt.VerifyJWTToken(token) 
		if err != nil || claim == nil {
			jerr, ok := err.(*jwt.ValidationError)
			if !ok {
				ecode = errcode.UnauthorizedTokenError
			} else {
				switch jerr.Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			}
		} else {
			c.Set("UUID", claim.UUID)
		}
	}

	if ecode != errcode.Success {
		response := app.NewGResponse(c)
		response.SendErrResponse(ecode)
		c.Abort()
		return
	}

	c.Next()
}
