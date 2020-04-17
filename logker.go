// Copyright (c) 2020 Higker
// Open Source: MIT License
// Author: higker <deen.job@qq.com>
package logker

/*

 */

// Build console Logger
// Level: logger Level
// Zone : logger Time Zone
// PS: We privatized some functions, structures and variables.
// Not affecting normal use. ^_^ Good luck~
func NewClog(lev level, zone logTimeZone) Logger {
	consoleLog := &console{
		logLevel: lev,
		timeZone: zone,
		tz:       nil,
	}
	consoleLog.initTime()
	return consoleLog
}
