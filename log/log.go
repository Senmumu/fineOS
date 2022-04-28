package log

import (
	"fmt"
	"os"
)

type LogLevel int8

const (
	LoglvlDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelNone
)

const (
	logLevelDev      = "FineOSLogLevel"
	logLevelDevDebug = "debug"
	logLevelDevInfo  = "info"
	logLevelDevWarn  = "warn"
	logLevelDevError = "error"
	logLevelDevNone  = "none"
	defaultLogLevel  = LogLevelError
)

var (
	Level              LogLevel
	ErrInvalidLogLevel = fmt.Errorf("Invalid log level")
)

func init() {
	level := os.Getenv("FineOSLogLevel")
}
