package log

import (
	"fmt"
	"log"
)

// Logger is a generic logging interface
type Logger interface {
	Log(v ...interface{})
	Logf(format string, v ...interface{})
}

var (
	DefaultLogger Logger = &NopLogger{}
)

func Log(v ...interface{}) {
	DefaultLogger.Log(v...)
}

func Logf(format string, v ...interface{}) {
	DefaultLogger.Logf(format, v...)
}

// NopLogger is a logger with no operation.
type NopLogger struct{}

// Log does nothing.
func (l *NopLogger) Log(v ...interface{}) {}

// Logf does nothing.
func (l *NopLogger) Logf(format string, v ...interface{}) {}

// StdLogger uses the standard log package as the logger.
type StdLogger struct{}

// Log uses the log.Output in standard log library.
func (l *StdLogger) Log(v ...interface{}) {
	log.Output(3, fmt.Sprintln(v...))
}

// Logf uses the log.Output in standard log library.
func (l *StdLogger) Logf(format string, v ...interface{}) {
	log.Output(3, fmt.Sprintf(format, v...))
}
