// Copyright 2020, The Go Authors. All rights reserved.
// Author: OnlyOneFace
// Date: 2020/7/24

package logger

import (
	"io"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	DebugLevel  = "debug"
	InfoLevel   = "info"
	WarnLevel   = "warn"
	ErrorLevel  = "error"
	DpanicLevel = "dpanic"
	PanicLevel  = "panic"
	FatalLevel  = "fatal"
)

type Encoder func(zapcore.EncoderConfig) zapcore.Encoder

func NewZap(level string, encoderFunc Encoder, w io.Writer, fields ...zap.Field) *zap.Logger {
	core := zapcore.NewCore(
		encoderFunc(newEncoderConfig()),
		zap.CombineWriteSyncers(zapcore.AddSync(w)),
		newLevel(level),
	).With(fields) //自带node 信息
	//大于error增加堆栈信息
	return zap.New(core).WithOptions(zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.DPanicLevel))
}

func newEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:    "message",
		LevelKey:      "level",
		TimeKey:       "time",
		NameKey:       "name",
		CallerKey:     "caller",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeTime: func(i time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(i.Local().Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
}

func newLevel(level string) zapcore.Level {
	var l zapcore.Level
	switch strings.ToUpper(level) {
	case DebugLevel:
		l = zap.DebugLevel
	case InfoLevel:
		l = zap.InfoLevel
	case WarnLevel:
		l = zap.WarnLevel
	case ErrorLevel:
		l = zap.ErrorLevel
	case DpanicLevel:
		l = zap.DPanicLevel
	case PanicLevel:
		l = zap.PanicLevel
	case FatalLevel:
		l = zap.FatalLevel
	default:
		l = zap.InfoLevel
	}
	return l
}
