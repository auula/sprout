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
	DEBUG level = iota
	INFO
	WARNING
	ERROR
)

// LogASingleFileConstant
const (
	_ = iota
	KB ByteSize = 1 << (10*iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)
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