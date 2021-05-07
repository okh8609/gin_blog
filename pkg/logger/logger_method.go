package logger

import (
	"context"
	"fmt"
)

// 紀錄檔分級輸出
func (l *MyLogger) Debug(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(DebugLevel, fmt.Sprint(v...))
}

func (l *MyLogger) Debugf(ctx context.Context, format string, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(DebugLevel, fmt.Sprintf(format, v...))
}

func (l *MyLogger) Info(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(InfoLevel, fmt.Sprint(v...))
}

func (l *MyLogger) Infof(ctx context.Context, format string, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(InfoLevel, fmt.Sprintf(format, v...))
}

func (l *MyLogger) Warn(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(WarnLevel, fmt.Sprint(v...))
}

func (l *MyLogger) Warnf(ctx context.Context, format string, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(WarnLevel, fmt.Sprintf(format, v...))
}

func (l *MyLogger) Error(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(ErrorLevel, fmt.Sprint(v...))
}

func (l *MyLogger) Errorf(ctx context.Context, format string, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(ErrorLevel, fmt.Sprintf(format, v...))
}

func (l *MyLogger) Fatal(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(FatalLevel, fmt.Sprint(v...))
}

func (l *MyLogger) Fatalf(ctx context.Context, format string, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(FatalLevel, fmt.Sprintf(format, v...))
}

func (l *MyLogger) Panic(ctx context.Context, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(PanicLevel, fmt.Sprint(v...))
}

func (l *MyLogger) Panicf(ctx context.Context, format string, v ...interface{}) {
	l.WithContext(ctx).WithTrace().Output(PanicLevel, fmt.Sprintf(format, v...))
}
