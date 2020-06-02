// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/4/30 - 8:41 下午

package logker

import "fmt"

// ByteSize
type ByteSize float64

// Custom logging level type.
type level int

// Logging level
const (
	// Debug level
	DEBUG level = iota
	// Info level
	INFO
	// Warning level
	WARNING
	// Error level
	ERROR
)

// ByteSize
const (
	// ByteSize Unit
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

// Log logger
var Log Logger

type Logger interface {
	// Info level log
	Info(value string, arg ...interface{})
	// Debug level log
	Debug(value string, arg ...interface{})
	// Error level log
	Error(value string, arg ...interface{})
	// Warning level log
	Warning(value string, arg ...interface{})
}

func Info(value string, arg ...interface{}) {
	Log.Info(value, arg)
}
func Debug(value string, arg ...interface{}) {
	Log.Info(value, arg)
}
func Error(value string, arg ...interface{}) {
	Log.Info(value, arg)
}
func Warning(value string, arg ...interface{}) {
	Log.Info(value, arg)
}

func init() {
	// Custom logging message template
	format := "{level} - Time {time}  - Position {position} - Message {message}"
	logger, err := NewClog(DEBUG, Shanghai, format, InitAsync(Qs3w))
	if err != nil {
		panic("init logger failed !!")
	}
	Log = logger
}

// TODO: Logging level to string
func (lev level) toStr() string {
	switch lev {
	case DEBUG:
		return "[ DEBUG ]"
	case INFO:
		return "[  INFO ]"
	case WARNING:
		return "[WARNING]"
	case ERROR:
		return "[ ERROR ]"
	default:
		// TheProgramWonTComeHereThroughTheCase~
		return fmt.Sprintf("[ UNWORKABLE ] 「Level to string fail:%d」\n", lev)
		//panic("Level to string fail :" + fmt.Sprintf("%d", lev))
	}
}
