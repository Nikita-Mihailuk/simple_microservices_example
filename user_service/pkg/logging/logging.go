package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var l *zap.Logger

func GetLogger() *zap.Logger {
	return l
}

func init() {
	consoleEncoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey:          "time",
		LevelKey:         "level",
		CallerKey:        "caller",
		MessageKey:       "msg",
		EncodeLevel:      zapcore.CapitalColorLevelEncoder,
		EncodeTime:       zapcore.ISO8601TimeEncoder,
		EncodeCaller:     zapcore.ShortCallerEncoder,
		ConsoleSeparator: " | ",
	})

	consoleWriter := zapcore.AddSync(os.Stdout)
	core := zapcore.NewCore(consoleEncoder, consoleWriter, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())

	l = logger
}
