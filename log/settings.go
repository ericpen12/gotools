package log

import (
	"github.com/ericpen12/gotools/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var logger *zap.Logger

func init() {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 设置日志记录中时间的格式
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	// 设置
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 生成打印到日志文件中的encoder
	//fileEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	// 将日志等级标识设置为大写并且有颜色
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	// 返回完整调用路径
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	// 生成打印到console的encoder
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	c := config.Log()

	level := zapcore.InfoLevel

	if c.Level == "debug" {
		level = zapcore.DebugLevel
	}

	core := zapcore.NewTee(
		// 同时向控制台和文件写日志， 生产环境记得把控制台写入去掉，日志记录的基本是Debug 及以上，生产环境记得改成Info
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level),
		//zapcore.NewCore(fileEncoder, zapcore.AddSync(&logConfig.logRoll), zapcore.DebugLevel),
	)
	// 返回调用栈
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(2))
}
