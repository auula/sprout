// Copyright (c) 2020 Higker
// Open Source: MIT License
// Author: higker <deen.job@qq.com>
package logker

/*

 */

// Build console Logger
// Level: logger Level
// Zone : logger Time Zone
func NewClog(level Level, zone LogTimeZone) Logger {
	consoleLog := &console{
		logLevel: level,
		timeZone: zone,
		tz:       nil,
	}
	consoleLog.initTime()
	return consoleLog
}
