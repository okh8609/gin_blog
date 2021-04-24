package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/okh8609/gin_blog/global"

	// 通用翻譯器(使用CLDR資料)
	ut "github.com/go-playground/universal-translator"

	// 從CLDR專案(Unicode通用語言環境資料儲存函數庫)產生的一組多語言環境
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"

	// validator的翻譯器
	"github.com/go-playground/validator/v10"
	en_tr "github.com/go-playground/validator/v10/translations/en"
	zh_tr "github.com/go-playground/validator/v10/translations/zh"
)

func TranslationMiddleware(c *gin.Context) {
	// 根據c的locale資訊  產生翻譯成locale語言的翻譯器
	locale := c.GetHeader("locale")
	global.MyLogger.Debugf(c, "#c.GetHeader(\"locale\"): %v.", locale) // debug logging
	ut2 := ut.New(en.New(), zh.New())
	trans, _ := ut2.GetTranslator(locale)

	c.Set("trans", trans) //將翻譯器儲存到全域的c中，以便後續翻譯時使用
	// 把 `gin的Validator` 轉成 `go-playground/validator`
	vt, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		switch locale {
		case "en":
			en_tr.RegisterDefaultTranslations(vt, trans) // 為validator註冊translator
		default:
			zh_tr.RegisterDefaultTranslations(vt, trans)
		}
	}

	c.Next()
}
