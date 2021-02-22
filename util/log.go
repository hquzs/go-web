package util

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

// ZeroLogger zerolog logger
type ZeroLogger struct {
	Log zerolog.Logger
}

func init() {
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05.999" // time.RFC3339Nano
	zerolog.MessageFieldName = "msg"
}

// NewZeroLog init zerolog with app name, package name
func NewZeroLog(app string, pkg string) *ZeroLogger {
	return &ZeroLogger{
		Log: zerolog.New(os.Stdout).With().Timestamp().Str("app", app).Str("pkg", pkg).Logger(),
	}
}

// Debugf logs a message at level Debug on the standard logger.
func (l *ZeroLogger) Debugf(format string, v ...interface{}) {
	l.Log.Debug().Msgf(format, v...)
}

// Debug logs a message at level debug
func (l *ZeroLogger) Debug(v ...interface{}) {
	l.Log.Debug().Msg(fmt.Sprint(v...))
}

// Infof logs a message at level Info on the standard logger.
func (l *ZeroLogger) Infof(format string, v ...interface{}) {
	l.Log.Info().Msgf(format, v...)
}

// Info logs a message at level Info
func (l *ZeroLogger) Info(v ...interface{}) {
	l.Log.Info().Msg(fmt.Sprint(v...))
}

// Warnf logs a message at level Warn on the standard logger.
func (l *ZeroLogger) Warnf(format string, v ...interface{}) {
	l.Log.Warn().Msgf(format, v...)
}

// Warn logs a message at level Warn
func (l *ZeroLogger) Warn(v ...interface{}) {
	l.Log.Warn().Msg(fmt.Sprint(v...))
}

// Warning is same as Warn
func (l *ZeroLogger) Warning(v ...interface{}) {
	l.Warn(v...)
}

// Warningf is same as Warnf
func (l *ZeroLogger) Warningf(format string, v ...interface{}) {
	l.Warnf(format, v...)
}

// Errorf logs a message at level Error on the standard logger.
func (l *ZeroLogger) Errorf(format string, v ...interface{}) {
	l.Log.Error().Msgf(format, v...)
}

// Error logs a message at level Error
func (l *ZeroLogger) Error(v ...interface{}) {
	l.Log.Error().Msg(fmt.Sprint(v...))
}

// Fatalf logs a message at level Fatal on the standard logger.
func (l *ZeroLogger) Fatalf(format string, v ...interface{}) {
	l.Log.Fatal().Msgf(format, v...)
}

// Fatal logs a message at level Fatal
func (l *ZeroLogger) Fatal(v ...interface{}) {
	l.Log.Fatal().Msg(fmt.Sprint(v...))
}

// Panic log a message at level Panic
func (l *ZeroLogger) Panic(v ...interface{}) {
	l.Log.Panic().Msg(fmt.Sprint(v...))
}

// Panicf logs a message at level Panic on the standard logger
func (l *ZeroLogger) Panicf(format string, v ...interface{}) {
	l.Log.Panic().Msgf(format, v...)
}

// SetLevel set zerolog level, it's global setting
func SetLevel(level string) {
	switch strings.ToLower(level) {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	default:
		fmt.Printf("Warning: unknown err level: %s", level)
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}
