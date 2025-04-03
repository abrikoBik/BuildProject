package util

import (
	"errors"
	"fmt"
	"log/slog"
	"strings"
)

// Sets the log level to one of these:
// - debug
// - warn
// - info
func LogLevel(level string) error {
	allowedLevels := map[string]slog.Level{
		"debug": slog.LevelDebug,
		"warn":  slog.LevelWarn,
		"info":  slog.LevelInfo,
	}
	for name, value := range allowedLevels {
		if strings.Compare(level, name) == 0 {
			slog.SetLogLoggerLevel(value)
			return nil
		}
	}
	return errors.New("unknown log level")
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

func Info(service string, format string, elements ...any) {
	var builder strings.Builder
	addServicePrefix(&builder, service)
	builder.WriteString(fmt.Sprintf(format, elements...))
	slog.Info(builder.String())
}

func Error(service string, err error) {
	var builder strings.Builder
	builder.Grow(errorLen)
	builder.WriteString("!!! ")
	addServicePrefix(&builder, service)
	builder.WriteString(err.Error())
	slog.Error(builder.String())
}

func Debug(service string, elements ...any) {
	var builder strings.Builder
	addServicePrefix(&builder, service)
	slog.Debug(builder.String(), elements...)
}
