package pipeline

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

func defaultLog(level LogSeverity, format string, a ...interface{}) {
	switch level {
	case LogError, LogFatal, LogPanic:
		s := ""
		if format == "" {
			s = fmt.Sprint(a...)
		} else {
			s = fmt.Sprintf(format, a...)
		}
		reportEvent(elError, 0, s)
	case LogWarning:
		s := ""
		if format == "" {
			s = fmt.Sprint(a...)
		} else {
			s = fmt.Sprintf(format, a...)
		}
		reportEvent(elWarning, 0, s)
	}
}

type eventType int16

const (
	elSuccess eventType = 0
	elError   eventType = 1
	elWarning eventType = 2
	elInfo    eventType = 4
)

var reportEvent = func() func(eventType eventType, eventID int32, msg string) {
	advAPI32 := syscall.MustLoadDLL("AdvAPI32.dll")
	registerEventSource := advAPI32.MustFindProc("RegisterEventSourceW")

	sourceName, _ := os.Executable()
	sourceNameUTF16, _ := syscall.UTF16PtrFromString(sourceName)
	handle, _, lastErr := registerEventSource.Call(uintptr(0), uintptr(unsafe.Pointer(sourceNameUTF16)))
	if lastErr == nil { // On error, logging is a no-op
		return func(eventType eventType, eventID int32, msg string) {}
	}
	reportEvent := advAPI32.MustFindProc("ReportEventW")
	return func(eventType eventType, eventID int32, msg string) {
		s, _ := syscall.UTF16PtrFromString(msg)
		_, _, _ = reportEvent.Call(
			uintptr(handle),             // HANDLE  hEventLog
			uintptr(eventType),          // WORD    wType
			uintptr(0),                  // WORD    wCategory
			uintptr(eventID),            // DWORD   dwEventID
			uintptr(0),                  // PSID    lpUserSid
			uintptr(1),                  // WORD    wNumStrings
			uintptr(0),                  // DWORD   dwDataSize
			uintptr(unsafe.Pointer(&s)), // LPCTSTR *lpStrings
			uintptr(0))                  // LPVOID  lpRawData
	}
}()
