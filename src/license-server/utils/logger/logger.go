package logger

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

const (
	LoggingNone    = iota
	LoggingError   = iota
	LoggingWarning = iota
	LoggingInfo    = iota
	LoggingDebug   = iota
)

var (
	Debug   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func LogOnError(err error) bool {
	if err != nil {
		Error.Println(err)
		return true
	}
	return false
}

func getLogLabel(severity string) string {
	return "[" + severity + "] "
}

func getLogWriter(level, targetLevel int) io.Writer {
	if level >= targetLevel {
		return os.Stdout
	}
	return ioutil.Discard
}

// TODO: Save the log to a file.
func initializeLoggers() {
	currLoggingLevel := LoggingDebug
	loggingFlag := log.LstdFlags | log.Lshortfile
	Debug = log.New(getLogWriter(currLoggingLevel, LoggingDebug), getLogLabel("DEBUG"), loggingFlag)
	Info = log.New(getLogWriter(currLoggingLevel, LoggingInfo), getLogLabel("INFO"), loggingFlag)
	Warning = log.New(getLogWriter(currLoggingLevel, LoggingWarning), getLogLabel("WARNING"), loggingFlag)
	Error = log.New(getLogWriter(currLoggingLevel, LoggingError), getLogLabel("ERROR"), loggingFlag)
}

func init() {
	initializeLoggers()
}
