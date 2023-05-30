package logging

import (
	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var rollingW = zapcore.AddSync(&lumberjack.Logger{
	Filename:   "tinydns.log",
	MaxSize:    500,
	MaxBackups: 3,
	MaxAge:     28,
})

var rollingC = zapcore.NewCore(
	zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig()),
	rollingW,
	zap.InfoLevel)

var stdoutW = zapcore.Lock(zapcore.AddSync(colorable.NewColorableStdout()))

var stdoutC = zapcore.NewCore(
	zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
	stdoutW,
	zap.InfoLevel)

var logger = zap.New(zapcore.NewTee(stdoutC, rollingC))

func GetLogger(name string) *zap.Logger {
	return logger.Named(name)
}

func GetSugar(name string) *zap.SugaredLogger {
	return GetLogger(name).Sugar()
}
