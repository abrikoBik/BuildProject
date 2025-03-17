package util

import (
	"fmt"
	"log/slog"
	"strings"
)

var builder strings.Builder

func addServicePrefix(service string) {
	builder.WriteRune('[')
	builder.WriteString(service)
	builder.WriteString("] ")
}

func Info(service string, format string, elements ...any) {
	addServicePrefix(service)
	builder.WriteString(fmt.Sprintf(format, elements...))
	slog.Info(builder.String())
	builder.Reset()
}

func Error(service string, err error) {
	builder.WriteString("!!! ")
	addServicePrefix(service)
	builder.WriteString(err.Error())
	slog.Error(builder.String())
	builder.Reset()
}

func Debug(service string, elements ...any) {
	addServicePrefix(service)
	slog.Debug(builder.String(), elements...)
	builder.Reset()
}
