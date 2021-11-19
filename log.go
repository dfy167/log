package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

func init() {
	hook := lumberjack.Logger{
		Filename:   "./logs/1.log", // 日志文件路径
		MaxSize:    128,            // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 30,             // 日志文件最多保存多少个备份
		MaxAge:     7,              // 文件最多保存多少天
		Compress:   true,           // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	//caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	//filed := zap.Fields(zap.String("serviceName", "serviceName"))
	// 构造日志
	// Logger = zap.New(core, development, filed)
	logger = zap.New(core, development)
}

// Debug Debug
func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

// Info Info
func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

// Warn Warn
func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

// Error Error
func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

//DPanic DPanic
func DPanic(msg string, fields ...zap.Field) {
	logger.DPanic(msg, fields...)
}

// Panic Panic
func Panic(msg string, fields ...zap.Field) {
	logger.Panic(msg, fields...)
}

// Fatal Fatal
func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

// Sync Sync
func Sync() error {
	err := logger.Sync()
	return err
}

func Map2fields(m map[string]interface{}) []zap.Field {
	fields := make([]zap.Field, 0, len(m))
	for k, v := range m {
		fields = append(fields, zap.Any(k, v))
	}
	return fields
}
