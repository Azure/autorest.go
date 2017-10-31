// +build !windows,!nacl,!plan9

package pipeline

import (
	"log"
	"log/syslog"
)

func defaultLog(severity LogSeverity, format string, a ...interface{}) {
	if defaultLogger == nil {
		return // Return fast if we failed to create the logger.
	}
	switch severity {
	case LogFatal:
		if format == "" {
			defaultLogger.Fatal(a...)
		} else {
			defaultLogger.Fatalf(format, a...)
		}
	case LogPanic:
		if format == "" {
			defaultLogger.Panic(a...)
		} else {
			defaultLogger.Panicf(format, a...)
		}
	case LogError, LogWarning:
		if format == "" {
			defaultLogger.Print(a...)
		} else {
			defaultLogger.Printf(format, a...)
		}
	default: // Do not log less severe entries
	}
}

var defaultLogger = func() *log.Logger {
	l, _ := syslog.NewLogger(syslog.LOG_USER|syslog.LOG_WARNING, log.LstdFlags)
	return l
}()
