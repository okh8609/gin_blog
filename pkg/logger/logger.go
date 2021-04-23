package logger

import (
	"context"
	"fmt"
	"io"
	"log"
	"runtime"

	"github.com/gin-gonic/gin"
)

// 宣告記錄檔的等級
type Level int8

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	PanicLevel
)

func (l *Level) String() string {
	switch *l {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	case PanicLevel:
		return "panic"
	}
	return ""
}

// 宣告用來存放紀錄檔共用欄位的型別 (可以存放任何東西的map；使用者需要自己知道每個欄位的型態)
type Field map[string]interface{}

// 實作自己的紀錄檔類別
type MyLogger struct {
	_logger *log.Logger
	contex  context.Context
	field   Field
	caller  []string // 存放呼叫者資訊的字串陣列
}

func NewLogger(w io.Writer, prefix string, flag int) *MyLogger {
	return &MyLogger{_logger: log.New(w, prefix, flag)}
}

func (l *MyLogger) clone() *MyLogger {
	l2 := *l
	return &l2
}

// 一些設定的Method
func (l *MyLogger) WithLevel(lv Level) *MyLogger {
	// 設定紀錄檔等級
	// --- pass ---
	return &MyLogger{}
}

func (l *MyLogger) WithContext(ct context.Context) *MyLogger {
	// 設定紀錄檔內容屬性
	l2 := l.clone()
	l2.contex = ct
	return l2
}

func (l *MyLogger) WithField(fd Field) *MyLogger {
	// 設定紀錄檔公用欄位
	l2 := l.clone()
	if l2.field == nil {
		l2.field = make(Field)
	}
	for k, v := range fd {
		l2.field[k] = v
	}
	return l2
}

func (l *MyLogger) WithCaller(skip int) *MyLogger {
	// 設定目前 `某一層呼叫` 的堆疊資訊
	// (程式計數器、檔案資訊、行號)

	l2 := l.clone()

	// runtime.Caller() 可以返回函數調用stack的某一層的程式計數器、檔案資訊、行號。
	// 0 代表當前函數， 1 代表上一層調用者... 以此類推。
	pc, file, line, ok := runtime.Caller(skip)

	if ok {
		f := runtime.FuncForPC(pc)                                           //runtime.FuncForPC()返回一個 *Func，描述包含給定程序計數器地址的函數，否則為nil。
		l2.caller = []string{fmt.Sprintf("%s: %d %s", file, line, f.Name())} // ?? 不是應該要用附加的嗎
	}

	return l2
}

func (l *MyLogger) WithCallerFrame() *MyLogger {
	// 設定目前 `整個呼叫` 的堆疊資訊
	// --- pass ---
	return &MyLogger{}
}

func (l *MyLogger) WithGinContext() *MyLogger {
	// 加入 *gin.Context 的資訊
	ginContext, ok := l.contex.(*gin.Context) // 型態轉換(Type assertions)成 *gin.Context
	if ok {                                   // 看看轉換有沒有成功
		return l.WithField(Field{
			"trace_id": ginContext.MustGet("X-Trace-ID"),
			"span_id":  ginContext.MustGet("X-Span-ID"),
		})
	} else {
		return l
	}
}
