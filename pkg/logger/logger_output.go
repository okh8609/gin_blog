package logger

// 紀錄檔格式化輸出
func (l *MyLogger) Output(level Level, msg string) {

}

func (l *MyLogger) JSONf(level Level, msg string) map[string]interface{} {
	return map[string]interface{}{}
}
