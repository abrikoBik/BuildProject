package handlers

import (
	"net/http"
	"runtime"
	"runtime/debug"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func DebugInfo(w http.ResponseWriter, r *http.Request) {
	config := zap.NewDevelopmentEncoderConfig()
	l := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		zapcore.AddSync(w),
		zapcore.DebugLevel,
	))
	defer l.Sync()
	l.Debug("DebugInfo",
		zap.String("Compiler", runtime.Compiler),
		zap.String("Version", runtime.Version()),
		zap.String("GOOS", runtime.GOOS),
		zap.String("GOARCH", runtime.GOARCH),
		zap.Int("GOMAXPROCS", runtime.GOMAXPROCS(0)),
		zap.Int("goroutines", runtime.NumGoroutine()),
		zap.Int("cpuCount", runtime.NumCPU()),
	)
}

func DebugGC(w http.ResponseWriter, r *http.Request) {
	config := zap.NewDevelopmentEncoderConfig()
	l := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		zapcore.AddSync(w),
		zapcore.DebugLevel,
	))
	defer l.Sync()
	var gcstats debug.GCStats
	debug.ReadGCStats(&gcstats)
	l.Debug("DebugGC",
		zap.Time("LastGC", gcstats.LastGC),
		zap.Int64("NumGC", gcstats.NumGC),
		zap.Durations("Pause", gcstats.Pause),
		zap.Times("PauseEnd", gcstats.PauseEnd),
		zap.Duration("PauseTotal", gcstats.PauseTotal),
	)
}
