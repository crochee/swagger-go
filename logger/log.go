// Copyright 2020, The Go Authors. All rights reserved.
// Author: OnlyOneFace
// Date: 2020/7/24

package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

var _logger *zap.Logger

// InitLog 日志初始化
func InitLog(level string) {
	_logger = NewZap(
		DebugLevel,
		zapcore.NewConsoleEncoder,
		os.Stdout)
}

// SetLoggerWriter
func SetLoggerWriter(writer io.Writer) {
	_logger = NewZap(
		DebugLevel,
		zapcore.NewConsoleEncoder,
		writer)
}

func Debugf(format string, args ...interface{}) {
	_logger.Sugar().Debugf(format, args...)
}

func Debug(format string) {
	_logger.Debug(format)
}

func Infof(format string, args ...interface{}) {
	_logger.Sugar().Infof(format, args...)
}

func Info(format string) {
	_logger.Info(format)
}

func Warnf(format string, args ...interface{}) {
	_logger.Sugar().Warnf(format, args...)
}

func Warn(format string) {
	_logger.Info(format)
}

func Errorf(format string, args ...interface{}) {
	_logger.Sugar().Errorf(format, args...)
}

func Error(format string) {
	_logger.Error(format)
}

func Panicf(format string, args ...interface{}) {
	_logger.Sugar().Panicf(format, args...)
	_ = _logger.Sync()
}

func Panic(format string) {
	_logger.Panic(format)
	_ = _logger.Sync()
}

func Fatalf(format string, args ...interface{}) {
	_logger.Sugar().Fatalf(format, args...)
	_ = _logger.Sync()
}

func Fatal(format string) {
	_logger.Fatal(format)
	_ = _logger.Sync()
}
