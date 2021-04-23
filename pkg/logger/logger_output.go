package logger

import (
	"encoding/json"
	"time"
)

// 紀錄檔格式化輸出
func (l *MyLogger) toJSON(level Level, msg string) map[string]interface{} {
	data := make(map[string]interface{})
	data["level"] = level.String()
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = msg
	data["caller"] = l.caller
	if len(l.field) > 0 {
		for k, v := range l.field {
			if _, ok := data[k]; !ok { // 如果 data[k] 不存在
				data[k] = v // 就插入新的值
			}
		}
	}

	return data
}

func (l *MyLogger) Output(level Level, msg string) {
	body, _ := json.Marshal(l.toJSON(level, msg))
	content := string(body)
	switch level {
	case DebugLevel:
		l._logger.Print(content)
	case InfoLevel:
		l._logger.Print(content)
	case WarnLevel:
		l._logger.Print(content)
	case ErrorLevel:
		l._logger.Print(content)
	case FatalLevel:
		l._logger.Fatal(content)
	case PanicLevel:
		l._logger.Panic(content)
	}
}
