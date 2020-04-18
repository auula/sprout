// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>

package logker

import "fmt"

// Custom logging level type.
type level int

const (
	DEBUG level = iota
	INFO
	WARNING
	ERROR
)
const (
	// Log file size
	// You can also customize. Unit is kb.
	MB10  int64 = 10 * 1024 * 1024
	MB100 int64 = 100 * 1024 * 1024
	GB1   int64 = 10 * MB100
)

// Type Log Logger
// Logger interface.
// Abstract Logger object.
type Logger interface {
	// English: Info level log
	Info(value string, arg ...interface{})
	// Debug level log
	Debug(value string, arg ...interface{})
	// Error level log
	Error(value string, arg ...interface{})
	// Warning level log
	Warning(value string, arg ...interface{})
}

// TODO: Logging level to string
func (lev level) toStr() string {
	switch lev {
	case DEBUG:
		return " DEBUG "
	case INFO:
		return "  INFO "
	case WARNING:
		return "WARNING"
	case ERROR:
		return " ERROR "
	default:
		panic("Level to string fail :" + fmt.Sprintf("%d", lev))
	}
}
