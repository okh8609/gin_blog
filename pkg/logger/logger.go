package logger

import (
	"context"
	"io"
	"log"
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
	caller  []string
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
	return &MyLogger{}
}

func (l *MyLogger) WithField(fd Field) *MyLogger {
	// 設定紀錄檔公用欄位
	return &MyLogger{}

}

func (l *MyLogger) WithContext(ct context.Context) *MyLogger {
	// 設定紀錄檔內容屬性
	return &MyLogger{}
}

func (l *MyLogger) WithCaller() *MyLogger {
	// 設定目前 `某一層呼叫` 的堆疊資訊
	// (程式計數器、檔案資訊、行號)
	return &MyLogger{}
}

func (l *MyLogger) WithCallerFrame() *MyLogger {
	// 設定目前 `整個呼叫` 的堆疊資訊
	return &MyLogger{}
}

func (l *MyLogger) WithTrace() *MyLogger {
	// 
	return &MyLogger{}
}

