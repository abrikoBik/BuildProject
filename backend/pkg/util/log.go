package util

import (
	"fmt"
	"strings"

	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	Service string
	Logger  zap.SugaredLogger
}

func CreateLogger(service string) Logger {
	config := zap.NewDevelopmentEncoderConfig()
	config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	l := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(config),
		zapcore.AddSync(colorable.NewColorableStdout()),
		zapcore.DebugLevel,
	))
	defer l.Sync()
	return Logger{
		Service: service,
		Logger:  *l.Sugar(),
	}
}

// +1 for an extra space
const (
	bracketLen = 2 + 1
	errorLen   = 3 + 1
)

func addServicePrefix(builder *strings.Builder, service string) {
	builder.Grow(len(service) + bracketLen)
	builder.WriteRune('[')
	builder.WriteString(service)
	builder.WriteString("] ")
}

func (l Logger) Info(format string, elements ...any) {
	var builder strings.Builder
	addServicePrefix(&builder, l.Service)
	builder.WriteString(fmt.Sprintf(format, elements...))
	l.Logger.Infof(builder.String())
}

func (l Logger) Error(err error) {
	var builder strings.Builder
	builder.Grow(errorLen)
	builder.WriteString("!!! ")
	addServicePrefix(&builder, l.Service)
	builder.WriteString(err.Error())
	l.Logger.Errorln(builder.String())
}

func (l Logger) Debug(elements ...any) {
	var builder strings.Builder
	addServicePrefix(&builder, l.Service)
	l.Logger.Debugw(builder.String(), elements...)
}
