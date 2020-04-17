// Copyright (c) 2020 Higker
// Open Source: MIT License
// Author: higker <deen.job@qq.com>
package logker

import "os"

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

func NewFlog(lev level, wheErr bool, zone logTimeZone, dir string, fileName string, power os.FileMode) *fileLog {
	fg := &fileLog{
		logLevel:  lev,
		wheError:  false,
		directory: dir,
		fileName:  fileName,
		file:      nil,
		errFile:   nil,
		tz:        nil,
		timeZone:  zone,
		power:     power,
	}
	fg.file, _ = fg.initFilePtr()
	return fg
}
