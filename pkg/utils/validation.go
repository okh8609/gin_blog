package utils
import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// 定義一些型別及方法

type ValidationError struct {
	Key string
	Msg string
}

func (e *ValidationError) Error() string {
	// implement: type error interface{}
	return e.Msg
}

type ValidationErrors []*ValidationError

func (es ValidationErrors) Errors() []string {
	var errs []string
	for _, e := range es {
		errs = append(errs, e.Error())
	}
	return errs
}

func (es ValidationErrors) Error() string {
	// implement: type error interface{}
	return strings.Join(es.Errors(), ", ")
}

// 主要函式
func BindAndValid(c *gin.Context, v interface{}) (bool, ValidationErrors) {
	// 回傳：驗證是否成功，驗證失敗的錯誤訊息
	// related with internal/middleware/translation.go
	e := c.ShouldBind(v) // 進行參數綁定和參數驗證
	if e != nil {
		var es ValidationErrors
		// fmt.Printf("#> %v", e)
		ves, ok := e.(validator.ValidationErrors)
		if !ok {
			return false, es // 空的
		}

		ut2 := ut.New(en.New())
		trans, _ := ut2.GetTranslator("")
		// trans, _ := (c.Value("trans")).(ut.Translator)
		for k, m := range ves.Translate(trans) {
			// fmt.Printf("#> k=%v, m=%v.\n", k, m)
			es = append(es, &ValidationError{
				Key: k,
				Msg: m,
			})
		}
		return false, es
	}
	return true, nil
}
