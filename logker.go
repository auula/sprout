// Copyright (c) 2020 Higker
// Open Source: MIT License
// Author: higker <deen.job@qq.com>
package logker

import (
	"os"
)

// This is package version
const Version = "1.1.3"

/*
 ____ ____ ____ ____ ____ ____
||L |||o |||g |||K |||e |||r ||
||__|||__|||__|||__|||__|||__||
|/__\|/__\|/__\|/__\|/__\|/__\|

zh_CN:
LogKer是Golang语言的日志操作库.
1.控制台输出 (已经实现)
2.文件输出 (已经实现)
3.WebSocket输出 (未来将会支持)
4.网络kafka输出  (未来将会支持)

*/

//type Formatting string

// Build console Logger
func NewClog(lev level, zone logTimeZone, formatting string) Logger {
	consoleLog := &console{
		// Level: logger Level
		logLevel: lev,
		// Zone : logger Time Zone
		timeZone: zone,
		tz:       nil,
		// Log Message Format Card
		formatting: "",
	}
	consoleLog.formatting = buildFormat(formatting)

	consoleLog.initTime()
	return consoleLog
}

// Build File logger
func NewFlog(lev level, wheErr bool, zone logTimeZone, dir string, fileName string, size int64, power os.FileMode, formatting string) Logger {
	fg := &fileLog{
		// logLevel:    lev,        logging level
		logLevel: lev,
		// wheError:    wheErr,     whether enable  error alone file
		wheError: wheErr,
		// directory:   dir,	    logging file save directory
		directory: dir,
		// fileName:    fileName,   logging save file name
		fileName: fileName,
		file:     nil,
		errFile:  nil,
		tz:       nil,
		// timeZone:    zone,	    load time zone format
		timeZone: zone,
		// power:       power,      file system power
		power: power,
		// fileMaxSize: size,       logging alone file max size
		fileMaxSize: size,
		// MessageMatchingCard
		formatting: "",
	}
	fg.formatting = buildFormat(formatting)
	fg.tz = &timeZone{TimeZoneStr: fg.timeZone}
	fg.file = fg.initFilePtr()
	if fg.isEnableErr() {
		fg.errFile = fg.initErrPtr()
	}
	return fg
}
