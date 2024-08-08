package log

import (
	"fmt"
	"github.com/ericpen12/gotools/config"
	"github.com/ericpen12/gotools/pkg"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var logger *zap.Logger

type Config struct {
	Level    string
	Filename string
}

func init() {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 设置日志记录中时间的格式
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	// 设置
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 生成打印到日志文件中的encoder
	fileEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	// 将日志等级标识设置为大写并且有颜色
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	// 返回完整调用路径
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	// 生成打印到console的encoder
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	var cfg Config
	_ = config.Load("log", &cfg)

	level := zapcore.InfoLevel
	if cfg.Level == "debug" {
		level = zapcore.DebugLevel
	}

	core := zapcore.NewTee(
		// 同时向控制台和文件写日志， 生产环境记得把控制台写入去掉，日志记录的基本是Debug 及以上，生产环境记得改成Info
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(writerSyncer(zapcore.InfoLevel, cfg)), zapcore.InfoLevel),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(writerSyncer(zapcore.WarnLevel, cfg)), zapcore.WarnLevel),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(writerSyncer(zapcore.ErrorLevel, cfg)), zapcore.ErrorLevel),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(writerSyncer(zapcore.FatalLevel, cfg)), zapcore.FatalLevel),
	)
	// 返回调用栈
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(2))
}

func writerSyncer(level zapcore.Level, cfg Config) zapcore.WriteSyncer {
	filename := cfg.Filename
	if filename == "" {
		homePath, _ := os.UserHomeDir()
		filename = fmt.Sprintf("%s/%s/%s.log", homePath, pkg.GetCurrentAppName(), level.String())
	}
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
		LocalTime:  true,
	}
	return zapcore.AddSync(lumberJackLogger)
}
