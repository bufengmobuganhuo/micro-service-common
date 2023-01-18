package common

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger *zap.SugaredLogger
)

func init() {
	// 日志文件名
	fileName := "micro.log"
	syncWriter := zapcore.AddSync(
		&lumberjack.Logger{
			// 文件名
			Filename: fileName,
			// 一个日志文件的最大大小
			MaxSize:    512,
			MaxBackups: 1,
			// 最大备份数
			LocalTime: true,
			Compress:  true,
		},
	)
	encoder := zap.NewProductionEncoderConfig()
	// 设置时间格式
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(zap.DebugLevel))

	log := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	logger = log.Sugar()
}

func Debug(args ...interface{}) {
	logger.Debug(args)
}

func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args)
}

func Info(args ...interface{}) {
	logger.Info(args)
}

func Error(args ...interface{}) {
	logger.Error(args)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args)
}
