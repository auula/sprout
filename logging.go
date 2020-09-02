package logker

import "time"

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

// Logging standard
type Logging interface {
	Info(args ...interface{})
	Error(args ...interface{})
	Debug(args ...interface{})
	Track(args ...interface{})
	Warning(args ...interface{})
}

// IOW  io standard
type IOW interface {
	Writer()
}

// LogItem logging message item
type LogItem struct {
	time     *time.Time
	Line     int64
	FuncName string
	FileName string
	Level    level
}
