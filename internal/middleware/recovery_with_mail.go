package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/okh8609/gin_blog/global"
	"github.com/okh8609/gin_blog/pkg/app"
	"github.com/okh8609/gin_blog/pkg/email"
	"github.com/okh8609/gin_blog/pkg/errcode"
)

func Recovery(c *gin.Context) {

	defailtMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.Email.Host,
		Port:     global.Email.Port,
		From:     global.Email.From,
		UserName: global.Email.UserName,
		Password: global.Email.Password,
		IsSSL:    global.Email.IsSSL,
	})

	defer func() {
		if err := recover(); err != nil {
			global.MyLogger.WithCallerFrame().Errorf(c, "panic recover err: %v", err)

			err := defailtMailer.SendMail(
				global.Email.To,
				fmt.Sprintf("異常拋出，發生時間: %d", time.Now().Unix()),
				fmt.Sprintf("錯誤信息: %v", err), // 收到信要去查log
			)
			if err != nil {
				global.MyLogger.Panicf(c, "mail.SendMail err: %v", err)
			}

			app.NewGResponse(c).SendErrResponse(errcode.ServerError)
			c.Abort()
		}
	}()

	c.Next()
}
