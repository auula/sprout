// Copyright (c) 2020 Higker
// Open Source: MIT License
// Author: higker <deen.job@qq.com>
package logker

/*

 */

// Build console Logger
// Level: logger Level
// Zone : logger Time Zone
func NewClog(lev level, zone LogTimeZone) Logger {
	consoleLog := &console{
		logLevel: lev,
		timeZone: zone,
		tz:       nil,
	}
	consoleLog.initTime()
	return consoleLog
}
