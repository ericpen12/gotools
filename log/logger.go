package log

import (
	"fmt"
	"go.uber.org/zap/zapcore"
)

func Info(data ...any) {
	do(zapcore.InfoLevel, "", data...)
}

func Debug(data ...any) {
	do(zapcore.DebugLevel, "", data...)
}

func Error(data ...any) {
	do(zapcore.ErrorLevel, "", data...)
}

func Warn(data ...any) {
	do(zapcore.WarnLevel, "", data...)
}

func Fatal(data ...any) {
	do(zapcore.FatalLevel, "", data...)
}

func Infof(template string, data ...any) {
	do(zapcore.InfoLevel, template, data...)
}

func Debugf(template string, data ...any) {
	do(zapcore.DebugLevel, template, data...)
}

func Errorf(template string, data ...any) {
	do(zapcore.ErrorLevel, template, data...)
}

func Warnf(template string, data ...any) {
	do(zapcore.WarnLevel, template, data...)
}

func Fatalf(template string, data ...any) {
	do(zapcore.FatalLevel, template, data...)
}

func formatMsg(template string, args ...any) string {
	if template != "" {
		return fmt.Sprintf(template+"", args...)
	}
	for i, arg := range args {
		if i <= len(args)-1 {
			args[i] = fmt.Sprintf("%v ", arg)
		}
	}
	return fmt.Sprint(args...)
}

func do(level zapcore.Level, template string, args ...any) {
	switch level {
	case zapcore.DebugLevel:
		logger.Debug(formatMsg(template, args...))
	case zapcore.InfoLevel:
		logger.Info(formatMsg(template, args...))
	case zapcore.WarnLevel:
		logger.Warn(formatMsg(template, args...))
	case zapcore.ErrorLevel:
		logger.Error(formatMsg(template, args...))
	case zapcore.FatalLevel:
		logger.Fatal(formatMsg(template, args...))
	}
}
